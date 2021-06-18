/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_DiscoveryAgentGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "DiscoveryAgent",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	DiscoveryAgentScope = "Environment"

	DiscoveryAgentResourceName = "discoveryagents"
)

func DiscoveryAgentGVK() apiv1.GroupVersionKind {
	return _DiscoveryAgentGVK
}

func init() {
	apiv1.RegisterGVK(_DiscoveryAgentGVK, DiscoveryAgentScope, DiscoveryAgentResourceName)
}

// DiscoveryAgent Resource
type DiscoveryAgent struct {
	apiv1.ResourceMeta

	Owner struct{} `json:"owner"`

	Spec DiscoveryAgentSpec `json:"spec"`

	Status DiscoveryAgentStatus `json:"status"`
}

// FromInstance converts a ResourceInstance to a DiscoveryAgent
func (res *DiscoveryAgent) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &DiscoveryAgentSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = DiscoveryAgent{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// DiscoveryAgentFromInstanceArray converts a []*ResourceInstance to a []*DiscoveryAgent
func DiscoveryAgentFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*DiscoveryAgent, error) {
	newArray := make([]*DiscoveryAgent, 0)
	for _, item := range fromArray {
		res := &DiscoveryAgent{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*DiscoveryAgent, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a DiscoveryAgent to a ResourceInstance
func (res *DiscoveryAgent) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = DiscoveryAgentGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
