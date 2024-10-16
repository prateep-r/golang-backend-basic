package product

import (
	"errors"
	"net/http"
	"training/app"
	"training/validator"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitEndpoints(r gin.IRoutes) {
	r.POST("/", h.Save)
	r.PUT("/:productId", h.Update)
	r.DELETE("/:productId", h.Delete)
	r.GET("/:productId", h.Get)
}

func (h *Handler) Get(c *gin.Context) {
	var req GetProductRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.Response{Message: err.Error()})
		return
	}

	if err := validator.Validate(c, req); err != nil {
		c.JSON(http.StatusBadRequest, app.Response{Message: err.Error()})
		return
	}

	product, err := h.service.GetById(c, uuid.MustParse(req.ProductId))
	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
			c.JSON(http.StatusNotFound, app.Response{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, app.Response{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, GetProductResponse{
		ProductId:   product.ProductId.String(),
		ProductName: product.ProductName,
		Price:       product.Price,
	})
}

func (h *Handler) Save(c *gin.Context) {

	var req SaveProductRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.Response{Message: err.Error()})
		return
	}

	if err := validator.Validate(c, req); err != nil {
		c.JSON(http.StatusBadRequest, app.Response{Message: err.Error()})
		return
	}

	payload, err := h.service.Save(c, SaveProductPayload{
		ProductName: req.ProductName,
		Price:       req.Price,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, app.Response{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}

func (h *Handler) Update(c *gin.Context) {
	var req UpdateProductRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.Response{Message: err.Error()})
		return
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.Response{Message: err.Error()})
		return
	}

	if err := validator.Validate(c, req); err != nil {
		c.JSON(http.StatusBadRequest, app.Response{Message: err.Error()})
		return
	}

	err := h.service.Update(c, UpdateProductPayload{
		ProductId:   uuid.MustParse(req.ProductId),
		ProductName: req.ProductName,
		Price:       req.Price,
	})
	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
			c.JSON(http.StatusNotFound, app.Response{Message: err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, app.Response{Message: err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) Delete(c *gin.Context) {
	var req DeleteProductRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.Response{Message: err.Error()})
		return
	}

	if err := validator.Validate(c, req); err != nil {
		c.JSON(http.StatusBadRequest, app.Response{Message: err.Error()})
		return
	}

	err := h.service.Delete(c, uuid.MustParse(req.ProductId))
	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
			c.JSON(http.StatusNotFound, app.Response{Message: err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, app.Response{Message: err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
