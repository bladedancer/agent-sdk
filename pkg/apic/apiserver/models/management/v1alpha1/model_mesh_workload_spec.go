/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// MeshWorkloadSpec struct for MeshWorkloadSpec
type MeshWorkloadSpec struct {
	Resources []string                `json:"resources,omitempty"`
	Labels    map[string]string       `json:"labels,omitempty"`
	Ports     []MeshWorkloadSpecPorts `json:"ports,omitempty"`
}
