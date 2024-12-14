package repos

import (
	"context"

	"github.com/udevs/lesson3/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductRepository interface {
	Create(ctx context.Context, product *models.Product) (*models.Product, error)

	FindByID(ctx context.Context, id primitive.ObjectID) (*models.Product, error)

	FindAll(ctx context.Context, page, limit int, search string) ([]*models.Product, error)

	Update(ctx context.Context, id primitive.ObjectID, product *models.Product) (*models.Product, error)

	Delete(ctx context.Context, id primitive.ObjectID) error

	Count(ctx context.Context, search string) (int64, error)
}
