/*
 * Swagger banking api simulation - OpenAPI 3.0
 *
 * Banking simulation application
 *
 * API version: 1.0.11
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package src

import (
	"net/http"

	"github.com/dopefresh/banking/golang/banking/src/di"
	"github.com/dopefresh/banking/golang/banking/src/handlers"
	"github.com/dopefresh/banking/golang/banking/src/middlewares"
	"github.com/dopefresh/banking/golang/banking/src/utils"
	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	logger := utils.ProvideLogger()
	container := di.Container{Log: logger}
	kafkaProducer := container.GetKafkaProducer()
	defer kafkaProducer.Close()
	kafkaConsumer := container.GetKafkaConsumer()
	defer kafkaConsumer.Close()

	kafkaService := container.GetKafkaTransactionService()
	clientHandler := container.GetClientHandler(kafkaService)
	cardCRUDHandler := container.GetCardHandler()
	transactionHandler := container.GetTransactionHandler()
	kafkaHandler := container.GetKafkaHandler()

	jwtService := container.GetJWTService()

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middlewares.AuthMiddleware(utils.ProvideLogger(), jwtService))

	for _, route := range GetRoutes(clientHandler, cardCRUDHandler, transactionHandler, kafkaHandler) {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func GetRoutes(clientHandler handlers.ClientHandler, cardCRUDHandler handlers.CardHandler, transactionHandler handlers.TransactionHandler, kafkaHandler handlers.KafkaHandler) Routes {
	return Routes{
		{
			"Index",
			http.MethodGet,
			"/",
			Index,
		},

		{
			"AddCard",
			http.MethodPost,
			"/api/v1/cards",
			cardCRUDHandler.AddCard,
		},

		{
			"GetCard",
			http.MethodGet,
			"/api/v1/cards/:number",
			cardCRUDHandler.GetCard,
		},

		{
			"UpdateCard",
			http.MethodPut,
			"/api/v1/cards/:number",
			cardCRUDHandler.UpdateCard,
		},

		{
			"DeleteCard",
			http.MethodDelete,
			"/api/v1/cards/:number",
			cardCRUDHandler.DeleteCard,
		},

		{
			"GetClient",
			http.MethodGet,
			"/api/v1/client",
			clientHandler.GetClient,
		},

		{
			"UpdateClient",
			http.MethodPut,
			"/api/v1/client",
			clientHandler.UpdateClient,
		},

		{
			"DepositMoney",
			http.MethodPost,
			"/api/v1/client/deposit/",
			clientHandler.DepositMoney,
		},

		{
			"WithdrawMoney",
			http.MethodPost,
			"/api/v1/client/withdraw/",
			clientHandler.WithdrawMoney,
		},

		{
			"TransferMoney",
			http.MethodPost,
			"/api/v1/client/transfer/",
			clientHandler.TransferMoney,
		},

		{
			"GetTransaction",
			http.MethodGet,
			"/api/v1/transactions/:transactionId/",
			transactionHandler.GetTransaction,
		},

		{
			"GetTransactions",
			http.MethodGet,
			"/api/v1/transactions/",
			transactionHandler.GetTransactions,
		},
		{
			"GetMessage",
			http.MethodGet,
			"/api/v1/kafka/readNextMessage/",
			kafkaHandler.GetNextMessage,
		},
	}
}
