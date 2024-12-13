package repos

import (
	"context"

	"github.com/udevs/lesson3/models"
)

type (
	OrdersRepository interface {
		Create(ctx context.Context, order *models.Order) (*models.Order, error)
		Get(ctx context.Context, id int) (*models.Order, error)
		Update(ctx context.Context, update *models.Order) (*models.Order, error)
		Delete(ctx context.Context, id int) (*models.Order, error)
		GetAll(ctx context.Context, page, limit int) ([]*models.Order, error)
	}
)
