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

	"github.com/dopefresh/banking/golang/banking/src/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CardHandler struct {
	CardService services.CardService
	Log         *zap.Logger
}

// AddCard - Add a new card
func (handler CardHandler) AddCard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteCard - delete card by number
func (handler CardHandler) DeleteCard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetCard - get card by number
func (handler CardHandler) GetCard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateCard - Update an existing card
func (handler CardHandler) UpdateCard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
