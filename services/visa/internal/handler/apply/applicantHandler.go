package apply

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marelinaa/visa-api/services/visa/internal/domain"
)

// Service defines the methods that a user service must implement.
type ApplyService interface {
	Apply(ctx context.Context, application domain.Application) error
}

type ApplyHandler struct {
	service ApplyService
}

func NewApplicantHandler(srv ApplyService) ApplyHandler {
	h := ApplyHandler{
		service: srv,
	}

	return h
}

// DefineRoutes defines routes for handling different API endpoints
func (h *ApplyHandler) DefineRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")

	currency := v1.Group("/visa")
	{
		currency.POST("/apply", h.Apply)
	}
}

// Apply ...
func (h *ApplyHandler) Apply(c *gin.Context) {
	var parsedApplication domain.Application //parsing application
	err := json.NewDecoder(c.Request.Body).Decode(&parsedApplication)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrDecodingReqBody.Error()})

		return
	}

	err = h.service.Apply(c, parsedApplication)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, gin.H{"error": "successfully applied for visa"})
}
