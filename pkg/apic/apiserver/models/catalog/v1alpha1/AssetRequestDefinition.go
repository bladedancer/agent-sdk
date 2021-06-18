/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_AssetRequestDefinitionGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "AssetRequestDefinition",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	AssetRequestDefinitionScope = "Asset"

	AssetRequestDefinitionResourceName = "assetrequestdefinitions"
)

func AssetRequestDefinitionGVK() apiv1.GroupVersionKind {
	return _AssetRequestDefinitionGVK
}

func init() {
	apiv1.RegisterGVK(_AssetRequestDefinitionGVK, AssetRequestDefinitionScope, AssetRequestDefinitionResourceName)
}

// AssetRequestDefinition Resource
type AssetRequestDefinition struct {
	apiv1.ResourceMeta

	Owner struct{} `json:"owner"`

	References struct{} `json:"references"`

	Spec AssetRequestDefinitionSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a AssetRequestDefinition
func (res *AssetRequestDefinition) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &AssetRequestDefinitionSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = AssetRequestDefinition{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a AssetRequestDefinition to a ResourceInstance
func (res *AssetRequestDefinition) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = AssetRequestDefinitionGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
