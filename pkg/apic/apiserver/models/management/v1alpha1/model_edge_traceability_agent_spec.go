/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// TODO: file was auto-generated with EdgeTraceabilityAgentSpecConfig and EdgeDiscoveryAgentSpecLogging

// EdgeTraceabilityAgentSpec struct for EdgeTraceabilityAgentSpec
type EdgeTraceabilityAgentSpec struct {
	// The name of the Axway Edge API Gateway dataplane associated to this agent
	Dataplane string `json:"dataplane"`
	// The name of the Traceability Agent associated to this agent
	TraceabilityAgent string                      `json:"traceabilityAgent,omitempty"`
	Config            TraceabilityAgentSpecConfig `json:"config"`
	// Config            EdgeTraceabilityAgentSpecConfig `json:"config"`
	Logging DiscoveryAgentSpecLogging `json:"logging,omitempty"`
	// Logging           EdgeDiscoveryAgentSpecLogging   `json:"logging,omitempty"`
}
