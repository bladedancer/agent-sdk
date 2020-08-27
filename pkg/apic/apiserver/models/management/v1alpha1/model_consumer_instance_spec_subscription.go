/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ConsumerInstanceSpecSubscription struct for ConsumerInstanceSpecSubscription
type ConsumerInstanceSpecSubscription struct {
	// Defines if subscriptions need to be manually approved.
	// GENERATE: The following code has been modified after after code generation
	// 	AutoSubscribe bool `json:"autoSubscribe,omitempty"`
	AutoSubscribe bool `json:"autoSubscribe"`
	// Defines if subscriptions are allowed on the Catalog Item.
	Enabled bool `json:"enabled,omitempty"`
	// The name of a ConsumerSubscriptionDefinition kind that defines the schema and possible webhooks to get invoked.
	SubscriptionDefinition string `json:"subscriptionDefinition,omitempty"`
}
