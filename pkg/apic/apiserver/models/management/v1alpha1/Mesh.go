/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_MeshGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "Mesh",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	MeshScope = ""

	MeshResourceName = "meshes"
)

func MeshGVK() apiv1.GroupVersionKind {
	return _MeshGVK
}

func init() {
	apiv1.RegisterGVK(_MeshGVK, MeshScope, MeshResourceName)
}

// Mesh Resource
type Mesh struct {
	apiv1.ResourceMeta

	Spec struct{} `json:"spec"`
}

// FromInstance converts a ResourceInstance to a Mesh
func (res *Mesh) FromInstance(ri *apiv1.ResourceInstance) error {
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

	*res = Mesh{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// MeshFromInstanceArray converts a []*ResourceInstance to a []*Mesh
func MeshFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*Mesh, error) {
	newArray := make([]*Mesh, 0)
	for _, item := range fromArray {
		res := &Mesh{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*Mesh, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a Mesh to a ResourceInstance
func (res *Mesh) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = MeshGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
