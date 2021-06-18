/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_ResourceGroupGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "definitions",
			Kind:  "ResourceGroup",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	ResourceGroupScope = ""

	ResourceGroupResourceName = "groups"
)

func ResourceGroupGVK() apiv1.GroupVersionKind {
	return _ResourceGroupGVK
}

func init() {
	apiv1.RegisterGVK(_ResourceGroupGVK, ResourceGroupScope, ResourceGroupResourceName)
}

// ResourceGroup Resource
type ResourceGroup struct {
	apiv1.ResourceMeta

	Owner struct{} `json:"owner"`

	Spec struct{} `json:"spec"`
}

// FromInstance converts a ResourceInstance to a ResourceGroup
func (res *ResourceGroup) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &struct{}{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = ResourceGroup{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// ResourceGroupFromInstanceArray converts a []*ResourceInstance to a []*ResourceGroup
func ResourceGroupFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*ResourceGroup, error) {
	newArray := make([]*ResourceGroup, 0)
	for _, item := range fromArray {
		res := &ResourceGroup{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*ResourceGroup, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a ResourceGroup to a ResourceInstance
func (res *ResourceGroup) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = ResourceGroupGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
