/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_DocumentGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "Document",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	DocumentScope = "Asset"

	DocumentResourceName = "documents"
)

func DocumentGVK() apiv1.GroupVersionKind {
	return _DocumentGVK
}

func init() {
	apiv1.RegisterGVK(_DocumentGVK, DocumentScope, DocumentResourceName)
}

// Document Resource
type Document struct {
	apiv1.ResourceMeta

	Owner struct{} `json:"owner"`

	Spec DocumentSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a Document
func (res *Document) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &DocumentSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = Document{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a Document to a ResourceInstance
func (res *Document) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = DocumentGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
