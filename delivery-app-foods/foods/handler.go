package foods

import (
	"net/http"

	"github.com/SantiagoBedoya/delivery-app-foods/utils/httperrors"
	"github.com/gin-gonic/gin"
)

// Handler defines interface for handlers
type Handler interface {
	Create(c *gin.Context)
	UpdateByID(c *gin.Context)
	DeleteByID(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
}

type handler struct {
	service Service
}

// NewHandler creates and implements Handler interface
func NewHandler(service Service) Handler {
	return &handler{service}
}

func (h *handler) Create(c *gin.Context) {
	var data Food
	if err := c.ShouldBindJSON(&data); err != nil {
		httpErr := httperrors.NewBadRequestError("invalid JSON body")
		c.JSON(httpErr.StatusCode, httpErr)
		return
	}
	food, err := h.service.Create(&data)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusCreated, food)
}
func (h *handler) UpdateByID(c *gin.Context) {
	var data Food
	if err := c.ShouldBindJSON(&data); err != nil {
		httpErr := httperrors.NewBadRequestError("invalid JSON body")
		c.JSON(httpErr.StatusCode, httpErr)
		return
	}
	if err := h.service.UpdateByID(c.Param("foodID"), &data); err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.Status(http.StatusNoContent)
}
func (h *handler) DeleteByID(c *gin.Context) {
	if err := h.service.DeleteByID(c.Param("foodID")); err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.Status(http.StatusNoContent)
}
func (h *handler) GetAll(c *gin.Context) {
	foods, err := h.service.GetAll()
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, foods)
}
func (h *handler) GetByID(c *gin.Context) {
	food, err := h.service.GetByID(c.Param("foodID"))
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, food)
}
