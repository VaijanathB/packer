/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type VmCreateDisk struct {
	Id string `json:"id,omitempty"`
	Size float32 `json:"size,omitempty"`
	Service string `json:"service,omitempty"`
	Name string `json:"name,omitempty"`
	Cloud string `json:"cloud,omitempty"`
}
