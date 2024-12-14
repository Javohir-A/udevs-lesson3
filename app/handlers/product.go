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

// ProductsHandler defines handlers for product operations.
type ProductsHandler struct {
	productsRepo repos.ProductRepository
	logger       *zap.Logger
}

// NewProductsHandler creates a new ProductsHandler instance.
func NewProductsHandler(repo repos.ProductRepository, logger *zap.Logger) *ProductsHandler {
	return &ProductsHandler{
		productsRepo: repo,
		logger:       logger,
	}
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Add a new product to the database
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      models.Product  true  "Product details"
// @Success      201      {object}  models.Product
// @Failure      400      {object}  gin.H
// @Failure      500      {object}  gin.H
// @Router       /products [post]
func (h *ProductsHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		h.logger.Error("Invalid input", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdProduct, err := h.productsRepo.Create(c.Request.Context(), &product)
	if err != nil {
		h.logger.Error("Failed to create product", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

// GetProductByID godoc
// @Summary      Get product by ID
// @Description  Retrieve product details by its ID
// @Tags         products
// @Produce      json
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  models.Product
// @Failure      404  {object}  gin.H
// @Failure      500  {object}  gin.H
// @Router       /products/{id} [get]
func (h *ProductsHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		h.logger.Error("Invalid product ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.productsRepo.FindByID(c.Request.Context(), objID)
	if err != nil {
		h.logger.Error("Product not found", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetAllProducts godoc
// @Summary      Get all products
// @Description  Retrieve all products with optional pagination and search
// @Tags         products
// @Produce      json
// @Param        page    query     int     false  "Page number"
// @Param        limit   query     int     false  "Page size"
// @Param        search  query     string  false  "Search query"
// @Success      200     {array}   models.Product
// @Failure      500     {object}  gin.H
// @Router       /products [get]
func (h *ProductsHandler) GetAllProducts(c *gin.Context) {
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

	products, err := h.productsRepo.FindAll(c.Request.Context(), int(pageInt), int(limitInt), search)
	if err != nil {
		h.logger.Error("Failed to retrieve products", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, products)
}
