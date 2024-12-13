package handlers

import (
	"github.com/udevs/lesson3/repos"
	"go.uber.org/zap"
)

type ProductsHandler struct {
	productsRepo repos.ProductRepository
	logger       *zap.Logger
}
