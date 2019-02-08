/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)

type Event struct {
	Id string `json:"_id,omitempty"`
	Name string `json:"name,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	Queued time.Time `json:"queued,omitempty"`
	State string `json:"state,omitempty"`
	Stage string `json:"stage,omitempty"`
	Resource EventResource `json:"resource,omitempty"`
}
