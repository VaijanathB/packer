/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// WebsiteCreateCredential struct for WebsiteCreateCredential
type WebsiteCreateCredential struct {
	Password    []JournalCreateCredentialPassword  `json:"password,omitempty"`
	Certificate []AgentCreateCredentialCertificate `json:"certificate,omitempty"`
}
