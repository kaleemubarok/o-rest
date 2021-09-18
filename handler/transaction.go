package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"o-rest/entity"
	"o-rest/service"
)

type errorResponse struct {
	err string `json:"error"`
}

type handler struct {
	s service.OrderService
}

func NewHandlerService(orderService service.OrderService) *handler {
	return &handler{s: orderService}
}

func (h *handler) GetOrder(c *gin.Context) {
	orders, err := h.s.GetOrders()
	if err != nil {
		response := errorResponse{"Error to get orders"}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := orders
	c.JSON(http.StatusOK, response)
}

func (h *handler) CreateOrder(c *gin.Context)  {
	var input entity.OrderRqResponse
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := errorResponse{"Failed to create order, check your request"}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newOrder, err := h.s.CreateOrder(input)
	if err != nil {
		response := errorResponse{"An error during create process"}
		c.JSON(http.StatusInternalServerError, response)
	}

	response := newOrder
	c.JSON(http.StatusOK, response)
}

func (h *handler) UpdateOrder(c *gin.Context)  {
	var input entity.OrderRqResponse
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := errorResponse{"Failed, please check your request"}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newOrder, err := h.s.UpdateOrder(input)
	if err != nil {
		response := errorResponse{"An error during update process"}
		c.JSON(http.StatusInternalServerError, response)
	}

	response := newOrder
	c.JSON(http.StatusOK, response)
}

func (h *handler) DeleteOrder(c *gin.Context)  {

	var input struct{
		ID   uint `uri:"id" binding:"required"`
	}
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := errorResponse{"Error to get order Id"}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order := entity.OrderRqResponse{OrderID: input.ID}
	err = h.s.DeleteOrder(order)
	if err != nil {
		response := errorResponse{"An error during delete process"}
		c.JSON(http.StatusInternalServerError, response)
	}

	c.JSON(http.StatusOK, nil)
}

