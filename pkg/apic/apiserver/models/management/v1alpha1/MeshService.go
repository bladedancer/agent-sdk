/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_MeshServiceGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "MeshService",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	MeshServiceScope = "Mesh"

	MeshServiceResourceName = "meshservices"
)

func MeshServiceGVK() apiv1.GroupVersionKind {
	return _MeshServiceGVK
}

func init() {
	apiv1.RegisterGVK(_MeshServiceGVK, MeshServiceScope, MeshServiceResourceName)
}

// MeshService Resource
type MeshService struct {
	apiv1.ResourceMeta

	Owner struct{} `json:"owner"`

	Spec MeshServiceSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a MeshService
func (res *MeshService) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &MeshServiceSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = MeshService{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// MeshServiceFromInstanceArray converts a []*ResourceInstance to a []*MeshService
func MeshServiceFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*MeshService, error) {
	newArray := make([]*MeshService, 0)
	for _, item := range fromArray {
		res := &MeshService{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*MeshService, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a MeshService to a ResourceInstance
func (res *MeshService) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = MeshServiceGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
