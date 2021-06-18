/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_AssetMappingGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "AssetMapping",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	AssetMappingScope = "Environment"

	AssetMappingResourceName = "assetmappings"
)

func AssetMappingGVK() apiv1.GroupVersionKind {
	return _AssetMappingGVK
}

func init() {
	apiv1.RegisterGVK(_AssetMappingGVK, AssetMappingScope, AssetMappingResourceName)
}

// AssetMapping Resource
type AssetMapping struct {
	apiv1.ResourceMeta

	Owner struct{} `json:"owner"`

	Spec AssetMappingSpec `json:"spec"`

	Status AssetMappingStatus `json:"status"`
}

// FromInstance converts a ResourceInstance to a AssetMapping
func (res *AssetMapping) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &AssetMappingSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = AssetMapping{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a AssetMapping to a ResourceInstance
func (res *AssetMapping) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = AssetMappingGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
