/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// GovernanceAgentSpecConfigAddress The listener address configuration
type GovernanceAgentSpecConfigAddress struct {
	// The listener interface
	Interface string `json:"interface"`
	// The listener exposed port
	Port int32 `json:"port"`
}
