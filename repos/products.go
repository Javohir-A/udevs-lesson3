package repos

import (
	"context"

	"github.com/udevs/lesson3/models"
)

type (
	ProductsRepository interface {
		Create(ctx context.Context, product *models.Product) (*models.Product, error)
		Get(ctx context.Context, id int) (*models.Product, error)
		Update(ctx context.Context, update *models.Product) (*models.Product, error)
		Delete(ctx context.Context, id int) (*models.Product, error)
		GetAll(ctx context.Context, page, limit int) ([]*models.Product, error)
	}
)
