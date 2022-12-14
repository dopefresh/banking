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

type TransactionHandler struct {
	Service services.TransactionService
	Log     *zap.Logger
}

// GetTransaction - get transaction by id
func (handler TransactionHandler) GetTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetTransactions - get client's transactions
func (handler TransactionHandler) GetTransactions(c *gin.Context) {
	userIdRaw, _ := c.Get("userId")
	userId, _ := userIdRaw.(int64)
	transactions, err := handler.Service.GetClientTransactions(userId)
	if err != nil {
		handler.Log.Error("Error occured", zap.Error(err))
		c.JSON(http.StatusInternalServerError, "Some error")
	}
	c.JSON(http.StatusOK, transactions)
}
