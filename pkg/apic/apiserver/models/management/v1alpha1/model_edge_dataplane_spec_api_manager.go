/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// EdgeDataplaneSpecApiManager Axway API Manager configuration.
type EdgeDataplaneSpecApiManager struct {
	// The host name where Axway API Manager is deployed
	Host string `json:"host,omitempty"`
	// The Axway API Manager admin port. Defaults to 8075
	Port int32 `json:"port,omitempty"`
	// Interval the agent will poll the Axway API Manager. Defaults to '30s' indicating 30 seconds The value is a sequence of number with unit suffix. Example '30s', '1m30s'. Valid units:  * 'h' - Hours  * 'm' - Minutes  * 's' - Seconds
	PollInterval string `json:"pollInterval,omitempty"`
}
