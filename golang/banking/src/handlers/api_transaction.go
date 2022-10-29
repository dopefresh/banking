/*
 * Swagger banking api simulation - OpenAPI 3.0
 *
 * Banking simulation application
 *
 * API version: 1.0.11
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTransaction - get transaction by id
func GetTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetTransactions - get client's transactions
func GetTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
