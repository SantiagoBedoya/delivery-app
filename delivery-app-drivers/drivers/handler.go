package drivers

import (
	"net/http"

	"github.com/SantiagoBedoya/delivery-app-drivers/utils/httperrors"
	"github.com/gin-gonic/gin"
)

// Handler defines interface for handlers
type Handler interface {
	SignIn(c *gin.Context)
	SignUp(c *gin.Context)
}

type handler struct {
	service Service
}

// NewHandler creates an instance of handlers and implements the interface
func NewHandler(service Service) Handler {
	return &handler{service}
}

func (h *handler) SignIn(c *gin.Context) {
	var data Driver
	if err := c.ShouldBindJSON(&data); err != nil {
		httpErr := httperrors.NewBadRequestError("invalid JSON body")
		c.JSON(httpErr.StatusCode, httpErr)
		return
	}
	accessToken, err := h.service.SignIn(&data)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *handler) SignUp(c *gin.Context) {
	var data Driver
	if err := c.ShouldBindJSON(&data); err != nil {
		httpErr := httperrors.NewBadRequestError("invalid JSON body")
		c.JSON(httpErr.StatusCode, httpErr)
		return
	}
	customer, err := h.service.SignUp(&data)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusCreated, customer)
}
