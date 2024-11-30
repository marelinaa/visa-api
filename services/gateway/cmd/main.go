package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/marelinaa/visa-api/services/gateway/internal/config"
	"github.com/marelinaa/visa-api/services/gateway/internal/handler"

	"log"
	"net/http"
)

func main() {
	cfg := config.Load()

	// users := map[string]string{
	// 	"user1": "12345",
	// 	"user2": "54321",
	// }
	// gatewayService := service.NewGatewayService(users)

	gatewayHandler := handler.NewGatewayHandler(gatewayService, cfg.AuthServiceURL, cfg.CurrencyServiceURL)

	router := gin.Default()
	gatewayHandler.DefineRoutes(router)

	apiPort := fmt.Sprintf(":%s", cfg.APIPort)
	log.Printf("Starting server on %s\n", apiPort)
	log.Fatal(http.ListenAndServe(apiPort, router))
}
