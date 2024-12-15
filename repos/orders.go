package repos

import (
	"context"

	"github.com/udevs/lesson3/models"
)

type OrderRepository interface {
	Create(ctx context.Context, order *models.Order) (*models.Order, error)

	FindByID(ctx context.Context, id string) (*models.Order, error)

	FindAll(ctx context.Context, page, limit int, status string) ([]*models.Order, error)

	Update(ctx context.Context, id string, order *models.Order) (*models.Order, error)

	Delete(ctx context.Context, id string) error

	GenerateReport(ctx context.Context, startDate, endDate string) ([]*models.Order, error)

	Count(ctx context.Context, status string) (int64, error)
}
