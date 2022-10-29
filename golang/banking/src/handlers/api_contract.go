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

// AddContract - Add a new contract
func AddContract(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetContract - Get client's contract by id
func GetContract(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetContracts - Get client's contracts
func GetContracts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
