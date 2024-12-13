package repos

import (
	"context"

	"github.com/udevs/lesson3/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderRepository interface {
	Create(ctx context.Context, order *models.Order) (*models.Order, error)

	FindByID(ctx context.Context, id primitive.ObjectID) (*models.Order, error)

	FindAll(ctx context.Context, page, limit int, status string) ([]*models.Order, error)

	Update(ctx context.Context, id primitive.ObjectID, order *models.Order) (*models.Order, error)

	Delete(ctx context.Context, id primitive.ObjectID) error

	GenerateReport(ctx context.Context, startDate, endDate string) ([]*models.Order, error)

	Count(ctx context.Context, status string) (int64, error)
}
