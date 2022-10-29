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

// AddCard - Add a new card
func AddCard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteCard - delete card by number
func DeleteCard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetCard - get card by number
func GetCard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateCard - Update an existing card
func UpdateCard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}