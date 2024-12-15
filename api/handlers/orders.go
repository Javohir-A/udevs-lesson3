package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/udevs/lesson3/models"
	"github.com/udevs/lesson3/repos"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type OrdersHandler struct {
	orderRepo repos.OrderRepository
	logger    *zap.Logger
}

func NewOrdersHandler(ordRepo repos.OrderRepository, log *zap.Logger) *OrdersHandler {
	return &OrdersHandler{
		orderRepo: ordRepo,
		logger:    log,
	}
}

// CreateOrder godoc
// @Summary      Create a new order
// @Description  Add a new order to the database
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        order  body      models.Order  true  "Order details"
// @Success      201    {object}  models.Order
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /orders [post]
func (h *OrdersHandler) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		h.logger.Error("Invalid input", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdOrder, err := h.orderRepo.Create(c.Request.Context(), &order)
	if err != nil {
		h.logger.Error("Failed to create order", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, createdOrder)
}

// GetOrderByID godoc
// @Summary      Get order by ID
// @Description  Retrieve order details by its ID
// @Tags         orders
// @Produce      json
// @Param        id   path      string  true  "Order ID"
// @Success      200  {object}  models.Order
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /orders/{id} [get]
func (h *OrdersHandler) GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		h.logger.Error("Invalid order ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := h.orderRepo.FindByID(c.Request.Context(), objID.Hex())
	if err != nil {
		h.logger.Error("Order not found", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// GetAllOrders godoc
// @Summary      Get all orders
// @Description  Retrieve all orders with optional pagination and search
// @Tags         orders
// @Produce      json
// @Param        page    query     int     false  "Page number"
// @Param        limit   query     int     false  "Page size"
// @Param        search  query     string  false  "Search query"
// @Success      200     {array}   models.Order
// @Failure      400     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /orders [get]
func (h *OrdersHandler) GetAllOrders(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	search := c.Query("search")

	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil || pageInt < 1 {
		h.logger.Error("Invalid page parameter", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil || limitInt < 1 {
		h.logger.Error("Invalid limit parameter", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	orders, err := h.orderRepo.FindAll(c.Request.Context(), int(pageInt), int(limitInt), search)
	if err != nil {
		h.logger.Error("Failed to retrieve orders", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GenerateReport godoc
// @Summary      Generate a report of orders within a date range
// @Description  Retrieve all orders between the specified start and end dates
// @Tags         orders
// @Produce      json
// @Param        startDate  query     string  true  "Start date in YYYY-MM-DD format"
// @Param        endDate    query     string  true  "End date in YYYY-MM-DD format"
// @Success      200        {array}   models.Order
// @Failure      400        {object}  map[string]string
// @Failure      500        {object}  map[string]string
// @Router       /orders/report [get]
func (h *OrdersHandler) GenerateReport(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	if startDate == "" || endDate == "" {
		h.logger.Error("Missing date parameters")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both startDate and endDate are required"})
		return
	}

	orders, err := h.orderRepo.GenerateReport(c.Request.Context(), startDate, endDate)
	if err != nil {
		h.logger.Error("Failed to generate report", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate report"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// UpdateOrder godoc
// @Summary      Update an existing order
// @Description  Update an order's details
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id     path      string       true  "Order ID"
// @Param        order  body      models.Order true  "Order details"
// @Success      200    {object}  models.Order
// @Failure      400    {object}  map[string]string
// @Failure      404    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /orders/{id} [put]
func (h *OrdersHandler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		h.logger.Error("Invalid order ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		h.logger.Error("Invalid input", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedOrder, err := h.orderRepo.Update(c.Request.Context(), objID.Hex(), &order)
	if err != nil {
		h.logger.Error("Failed to update order", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(http.StatusOK, updatedOrder)
}

// DeleteOrder godoc
// @Summary      Delete an order
// @Description  Remove an order by its ID
// @Tags         orders
// @Produce      json
// @Param        id  path      string  true  "Order ID"
// @Success      200  {object} map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /orders/{id} [delete]
func (h *OrdersHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		h.logger.Error("Invalid order ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	if err := h.orderRepo.Delete(c.Request.Context(), objID.Hex()); err != nil {
		h.logger.Error("Failed to delete order", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
