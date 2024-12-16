package auth

import (
	"context"
	"net/http"

	"github.com/marelinaa/visa-api/services/visa/internal/domain"
	"github.com/marelinaa/visa-api/services/visa/internal/service"

	"github.com/gin-gonic/gin"
)

// Service defines the methods that a user service must implement.
type AuthService interface {
	SignUp(ctx context.Context, userInput domain.SignUpInput) error
	SignIn(ctx context.Context, userSignInInput domain.SignInInput) (service.Token, error)
}

// Handler manages HTTP requests and responses for user operations.
type AuthHandler struct {
	service AuthService
}

// NewHandler creates a new Handler with the given service.
func NewAuthHandler(srv AuthService) AuthHandler {
	h := AuthHandler{
		service: srv,
	}

	return h
}

// SignUp ...
func (h *AuthHandler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	var req domain.SignUpInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrDecodingReqBody.Error()})

		return
	}

	err := h.service.SignUp(ctx, req)
	if err != nil {
		//todo: return right http code
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, gin.H{"succesfully created profile": "go to Sign In page"})
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	ctx := c.Request.Context()

	var req domain.SignInInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrDecodingReqBody.Error()})

		return
	}

	token, err := h.service.SignIn(ctx, req)
	if err != nil {
		if err == domain.ErrUnauthorized {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, token)
}
