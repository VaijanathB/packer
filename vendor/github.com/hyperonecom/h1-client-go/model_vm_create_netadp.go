/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type VmCreateNetadp struct {
	Service string `json:"service,omitempty"`
	Network string `json:"network,omitempty"`
	Ip []string `json:"ip,omitempty"`
}
