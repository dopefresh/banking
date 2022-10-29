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

type Transaction struct {

	TransactionId int64 `json:"transactionId,omitempty"`

	TransactionType string `json:"transactionType,omitempty"`

	CardIdFrom string `json:"cardIdFrom,omitempty"`

	CardIdTo string `json:"cardIdTo,omitempty"`

	CardId string `json:"cardId,omitempty"`

	Summ float64 `json:"summ,omitempty"`

	TransactionDatetime string `json:"transactionDatetime,omitempty"`
}