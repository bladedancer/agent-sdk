/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_APISpecGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "APISpec",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	APISpecScope = "K8SCluster"

	APISpecResourceName = "apispecs"
)

func APISpecGVK() apiv1.GroupVersionKind {
	return _APISpecGVK
}

func init() {
	apiv1.RegisterGVK(_APISpecGVK, APISpecScope, APISpecResourceName)
}

// APISpec Resource
type APISpec struct {
	apiv1.ResourceMeta

	Owner struct{} `json:"owner"`

	Spec ApiSpecSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a APISpec
func (res *APISpec) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &ApiSpecSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = APISpec{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// APISpecFromInstanceArray converts a []*ResourceInstance to a []*APISpec
func APISpecFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*APISpec, error) {
	newArray := make([]*APISpec, 0)
	for _, item := range fromArray {
		res := &APISpec{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*APISpec, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a APISpec to a ResourceInstance
func (res *APISpec) AsInstance() (*apiv1.ResourceInstance, error) {
	m, err := json.Marshal(res.Spec)
	if err != nil {
		return nil, err
	}

	spec := map[string]interface{}{}
	err = json.Unmarshal(m, &spec)
	if err != nil {
		return nil, err
	}

	meta := res.ResourceMeta
	meta.GroupVersionKind = APISpecGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
