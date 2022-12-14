/*
 * Swagger banking api simulation - OpenAPI 3.0
 *
 * Banking simulation application
 *
 * API version: 1.0.11
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ContractCreate struct {
	ContractType string `json:"contractType,omitempty"`

	ConclusionDate string `json:"conclusionDate,omitempty"`

	ContractContent string `json:"contractContent,omitempty"`

	ClientIt int64 `json:"clientIt,omitempty"`
}
