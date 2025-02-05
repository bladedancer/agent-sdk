package agent

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	coreapi "github.com/Axway/agent-sdk/pkg/api"
	"github.com/Axway/agent-sdk/pkg/apic"
	apiV1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
	"github.com/Axway/agent-sdk/pkg/cache"
	"github.com/Axway/agent-sdk/pkg/config"
	"github.com/Axway/agent-sdk/pkg/jobs"
	utilErrors "github.com/Axway/agent-sdk/pkg/util/errors"
	hc "github.com/Axway/agent-sdk/pkg/util/healthcheck"
	"github.com/Axway/agent-sdk/pkg/util/log"
)

const (
	apiServerPageSize        = 100
	healthcheckEndpoint      = "central"
	attributesQueryParam     = "attributes."
	apiServerFields          = "name,title,attributes"
	serviceInstanceCache     = "ServiceInstances"
	serviceInstanceNameCache = "ServiceInstanceNames"
)

var discoveryCacheLock *sync.Mutex

func init() {
	discoveryCacheLock = &sync.Mutex{}
}

type discoveryCache struct {
	jobs.Job
	lastServiceTime  time.Time
	lastInstanceTime time.Time
	refreshAll       bool
}

func newDiscoveryCache(getAll bool) *discoveryCache {
	return &discoveryCache{
		lastServiceTime:  time.Time{},
		lastInstanceTime: time.Time{},
		refreshAll:       getAll,
	}
}

//Ready -
func (j *discoveryCache) Ready() bool {
	status := hc.GetStatus(healthcheckEndpoint)
	return status == hc.OK
}

//Status -
func (j *discoveryCache) Status() error {
	status := hc.GetStatus(healthcheckEndpoint)
	if status == hc.OK {
		return nil
	}
	return fmt.Errorf("could not establish a connection to APIC to update the cache")
}

//Execute -
func (j *discoveryCache) Execute() error {
	discoveryCacheLock.Lock()
	defer discoveryCacheLock.Unlock()
	log.Trace("executing API cache update job")
	j.updateAPICache()
	if agent.cfg.GetAgentType() == config.DiscoveryAgent {
		j.validateAPIServiceInstances()
	}
	fetchConfig()
	return nil
}

func (j *discoveryCache) updateAPICache() {
	log.Trace("updating API cache")

	// Update cache with published resources
	existingAPIs := make(map[string]bool)
	query := map[string]string{
		apic.FieldsKey: apiServerFields,
	}

	if !j.lastServiceTime.IsZero() && !j.refreshAll {
		query[apic.QueryKey] = fmt.Sprintf("%s>\"%s\"", apic.CreateTimestampQueryKey, j.lastServiceTime.Format(v1.APIServerTimeFormat))
	}
	apiServices, _ := GetCentralClient().GetAPIV1ResourceInstancesWithPageSize(query, agent.cfg.GetServicesURL(), apiServerPageSize)

	for _, apiService := range apiServices {
		if _, valid := apiService.Attributes[apic.AttrExternalAPIID]; !valid {
			continue // skip service without external api id
		}
		// Update the lastServiceTime based on the newest service found
		thisTime := time.Time(apiService.Metadata.Audit.CreateTimestamp)
		if j.lastServiceTime.Before(thisTime) {
			j.lastServiceTime = thisTime
		}

		externalAPIID := addItemToAPICache(*apiService)
		if externalAPIPrimaryKey, found := apiService.Attributes[apic.AttrExternalAPIPrimaryKey]; found {
			existingAPIs[externalAPIPrimaryKey] = true
		} else {
			existingAPIs[externalAPIID] = true
		}
	}

	if j.refreshAll {
		// Remove items that are not published as Resources
		cacheKeys := agent.apiMap.GetKeys()
		for _, key := range cacheKeys {
			if _, ok := existingAPIs[key]; !ok {
				agent.apiMap.Delete(key)
			}
		}
	}
}

func (j *discoveryCache) validateAPIServiceInstances() {
	if agent.apiValidator == nil {
		return
	}

	query := map[string]string{
		apic.FieldsKey: apiServerFields,
	}

	if !j.lastInstanceTime.IsZero() && !j.refreshAll {
		query[apic.QueryKey] = fmt.Sprintf("%s>\"%s\"", apic.CreateTimestampQueryKey, j.lastServiceTime.Format(v1.APIServerTimeFormat))
	}

	j.lastInstanceTime = time.Now()
	serviceInstances, err := GetCentralClient().GetAPIV1ResourceInstancesWithPageSize(query, agent.cfg.GetInstancesURL(), apiServerPageSize)
	if err != nil {
		log.Error(utilErrors.Wrap(ErrUnableToGetAPIV1Resources, err.Error()).FormatError("APIServiceInstances"))
		return
	}

	for _, instance := range serviceInstances {
		if j.refreshAll {
			break // no need to do this loop when refreshing the entire cache
		}
		if _, valid := instance.Attributes[apic.AttrExternalAPIID]; !valid {
			continue // skip instance without external api id
		}
		// Update the lastInstanceTime based on the newest instance found
		thisTime := time.Time(instance.Metadata.Audit.CreateTimestamp)
		if j.lastInstanceTime.Before(thisTime) {
			j.lastInstanceTime = thisTime
		}
	}

	// When reloading all api service instances we can just write over the existing cache
	if !j.refreshAll {
		serviceInstances = j.loadServiceInstancesFromCache(serviceInstances)
	}
	serviceInstances = validateAPIOnDataplane(serviceInstances)
	j.saveServiceInstancesToCache(serviceInstances)
}

func (j *discoveryCache) saveServiceInstancesToCache(serviceInstances []*apiV1.ResourceInstance) {
	// Save all the instance names to make sure the map is unique
	instanceNames := make(map[string]struct{})
	for _, instance := range serviceInstances {
		instanceNames[instance.Name] = struct{}{}
	}
	cache.GetCache().Set(serviceInstanceCache, serviceInstances)
	cache.GetCache().Set(serviceInstanceNameCache, instanceNames)
}

func (j *discoveryCache) loadServiceInstancesFromCache(serviceInstances []*apiV1.ResourceInstance) []*apiV1.ResourceInstance {
	cachedInstancesInterface, err := cache.GetCache().Get(serviceInstanceCache)
	if err != nil {
		return serviceInstances
	}
	cachedInstancesNames, err := cache.GetCache().Get(serviceInstanceNameCache)
	if err != nil {
		return serviceInstances
	}
	cachedInstances := cachedInstancesInterface.([]*apiV1.ResourceInstance)
	for _, instance := range serviceInstances {
		// validate that the instance is not already in the array
		if _, found := cachedInstancesNames.(map[string]struct{})[instance.Name]; !found {
			cachedInstances = append(cachedInstances, instance)
		}
	}

	// return the full list
	return cachedInstances
}

var updateCacheForExternalAPIPrimaryKey = func(externalAPIPrimaryKey string) (interface{}, error) {
	query := map[string]string{
		apic.QueryKey: attributesQueryParam + apic.AttrExternalAPIPrimaryKey + "==\"" + externalAPIPrimaryKey + "\"",
	}

	return updateCacheForExternalAPI(query)
}

var updateCacheForExternalAPIID = func(externalAPIID string) (interface{}, error) {
	query := map[string]string{
		apic.QueryKey: attributesQueryParam + apic.AttrExternalAPIID + "==\"" + externalAPIID + "\"",
	}

	return updateCacheForExternalAPI(query)
}

var updateCacheForExternalAPIName = func(externalAPIName string) (interface{}, error) {
	query := map[string]string{
		apic.QueryKey: attributesQueryParam + apic.AttrExternalAPIName + "==\"" + externalAPIName + "\"",
	}

	return updateCacheForExternalAPI(query)
}

var updateCacheForExternalAPI = func(query map[string]string) (interface{}, error) {
	apiServerURL := agent.cfg.GetServicesURL()

	response, err := agent.apicClient.ExecuteAPI(coreapi.GET, apiServerURL, query, nil)
	if err != nil {
		return nil, err
	}
	apiService := apiV1.ResourceInstance{}
	json.Unmarshal(response, &apiService)
	addItemToAPICache(apiService)
	return apiService, nil
}

func validateAPIOnDataplane(serviceInstances []*apiV1.ResourceInstance) []*apiV1.ResourceInstance {
	cleanServiceInstances := make([]*apiV1.ResourceInstance, 0)
	// Validate the API on dataplane.  If API is not valid, mark the consumer instance as "DELETED"
	for _, serviceInstanceResource := range serviceInstances {
		if _, valid := serviceInstanceResource.Attributes[apic.AttrExternalAPIID]; !valid {
			continue // skip service instances without external api id
		}
		serviceInstance := &v1alpha1.APIServiceInstance{}
		serviceInstance.FromInstance(serviceInstanceResource)
		externalAPIID := serviceInstance.Attributes[apic.AttrExternalAPIID]
		externalAPIStage := serviceInstance.Attributes[apic.AttrExternalAPIStage]
		// Check if the consumer instance was published by agent, i.e. following attributes are set
		// - externalAPIID should not be empty
		// - externalAPIStage could be empty for dataplanes that do not support it
		if externalAPIID != "" && !agent.apiValidator(externalAPIID, externalAPIStage) {
			deleteServiceInstanceOrService(serviceInstance, externalAPIID, externalAPIStage)
		} else {
			cleanServiceInstances = append(cleanServiceInstances, serviceInstanceResource)
		}
	}
	return cleanServiceInstances
}

func shouldDeleteService(apiID, stage string) bool {
	// no agent-specific validator means to delete the service
	if agent.deleteServiceValidator == nil {
		return true
	}
	// let the agent decide if service should be deleted
	return agent.deleteServiceValidator(apiID, stage)
}

func deleteServiceInstanceOrService(serviceInstance *v1alpha1.APIServiceInstance, externalAPIID, externalAPIStage string) {
	if shouldDeleteService(externalAPIID, externalAPIStage) {
		log.Infof("API no longer exists on the dataplane; deleting the API Service and corresponding catalog item %s", serviceInstance.Title)
		// deleting the service will delete all associated resources, including the consumerInstance
		err := agent.apicClient.DeleteServiceByAPIID(externalAPIID)
		if err != nil {
			log.Error(utilErrors.Wrap(ErrDeletingService, err.Error()).FormatError(serviceInstance.Title))
		} else {
			log.Debugf("Deleted API Service for catalog item %s from Amplify Central", serviceInstance.Title)
		}
	} else {
		log.Infof("API no longer exists on the dataplane, deleting the catalog item %s", serviceInstance.Title)
		err := agent.apicClient.DeleteAPIServiceInstance(serviceInstance.Name)
		if err != nil {
			log.Error(utilErrors.Wrap(ErrDeletingCatalogItem, err.Error()).FormatError(serviceInstance.Title))
		} else {
			log.Debugf("Deleted catalog item %s from Amplify Central", serviceInstance.Title)
		}
	}
}

func addItemToAPICache(apiService apiV1.ResourceInstance) string {
	externalAPIID, ok := apiService.Attributes[apic.AttrExternalAPIID]
	if ok {
		externalAPIName := apiService.Attributes[apic.AttrExternalAPIName]
		if externalAPIPrimaryKey, found := apiService.Attributes[apic.AttrExternalAPIPrimaryKey]; found {
			// Verify secondary key and validate if we need to remove it from the apiMap (cache)
			if _, err := agent.apiMap.Get(externalAPIID); err != nil {
				agent.apiMap.Delete(externalAPIID)
			}

			agent.apiMap.SetWithSecondaryKey(externalAPIPrimaryKey, externalAPIID, apiService)
			agent.apiMap.SetSecondaryKey(externalAPIPrimaryKey, externalAPIName)
		} else {
			agent.apiMap.SetWithSecondaryKey(externalAPIID, externalAPIName, apiService)
		}
		log.Tracef("added api name: %s, id %s to API cache", externalAPIName, externalAPIID)
	}
	return externalAPIID
}
