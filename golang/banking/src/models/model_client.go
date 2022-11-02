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

type Client struct {
	ClientId int64 `json:"clientId,omitempty"`

	// Client first name
	FirstName string `json:"firstName,omitempty"`

	// Client last name
	LastName string `json:"lastName,omitempty"`

	Phone string `json:"phone,omitempty"`

	// Client registration address
	RegistrationAddress string `json:"registrationAddress,omitempty"`

	ResidentialAddress string `json:"residentialAddress,omitempty"`

	ClientType string `json:"clientType,omitempty"`

	// Main State registration number
	Ogrn string `json:"ogrn,omitempty"`

	// unique taxpayer identification number
	Inn string `json:"inn,omitempty"`

	// complements the inn
	Kpp string `json:"kpp,omitempty"`
}
