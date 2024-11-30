package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marelinaa/visa-api/services/visa/internal/domain"
	"github.com/marelinaa/visa-api/services/visa/internal/service"
)

type ApplicantHandler struct {
	service *service.ApplicantService
}

func NewApplicantHandler(srv *service.ApplicantService) *ApplicantHandler {
	return &ApplicantHandler{
		service: srv,
	}
}

// DefineRoutes defines routes for handling different API endpoints
func (h *ApplicantHandler) DefineRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")

	currency := v1.Group("/visa")
	{
		currency.POST("/apply", h.Apply)
	}
}

// Apply ...
func (h *ApplicantHandler) Apply(c *gin.Context) {
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
