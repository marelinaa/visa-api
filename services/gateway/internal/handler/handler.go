package handler

import (
	"github.com/marelinaa/visa-api/services/gateway/internal/service"

	"github.com/gin-gonic/gin"
)

type GatewayHandler struct {
	service        *service.GatewayService
	authServiceURL string
	visaServiceURL string
}

func NewGatewayHandler(srv *service.GatewayService, authServiceURL, visaServiceURL string) *GatewayHandler {
	return &GatewayHandler{
		service:        srv,
		authServiceURL: authServiceURL,
		visaServiceURL: visaServiceURL,
	}
}

// DefineRoutes defines routes for handling different API endpoints
func (h *GatewayHandler) DefineRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET("/sign-in", h.SignIn)

	currency := v1.Group("/currency", h.Authorize())
	{
		currency.GET("/date", h.GetCurrencyByDate)
		currency.GET("/history", h.GetCurrencyHistory)
	}
}
