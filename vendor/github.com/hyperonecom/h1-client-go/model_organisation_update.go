/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// OrganisationUpdate struct for OrganisationUpdate
type OrganisationUpdate struct {
	Name    string                    `json:"name,omitempty"`
	Billing OrganisationUpdateBilling `json:"billing,omitempty"`
}
