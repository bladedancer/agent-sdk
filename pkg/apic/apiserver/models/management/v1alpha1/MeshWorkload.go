/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_MeshWorkloadGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "MeshWorkload",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	MeshWorkloadScope = "Mesh"

	MeshWorkloadResourceName = "meshworkloads"
)

func MeshWorkloadGVK() apiv1.GroupVersionKind {
	return _MeshWorkloadGVK
}

func init() {
	apiv1.RegisterGVK(_MeshWorkloadGVK, MeshWorkloadScope, MeshWorkloadResourceName)
}

// MeshWorkload Resource
type MeshWorkload struct {
	apiv1.ResourceMeta

	Spec MeshWorkloadSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a MeshWorkload
func (res *MeshWorkload) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &MeshWorkloadSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = MeshWorkload{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// MeshWorkloadFromInstanceArray converts a []*ResourceInstance to a []*MeshWorkload
func MeshWorkloadFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*MeshWorkload, error) {
	newArray := make([]*MeshWorkload, 0)
	for _, item := range fromArray {
		res := &MeshWorkload{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*MeshWorkload, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a MeshWorkload to a ResourceInstance
func (res *MeshWorkload) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = MeshWorkloadGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
