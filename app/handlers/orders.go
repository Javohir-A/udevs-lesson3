package handlers

import (
	"github.com/udevs/lesson3/repos"
	"go.uber.org/zap"
)

type OrdersHandler struct {
	orderRepo repos.OrderRepository
	logger    *zap.Logger
}
