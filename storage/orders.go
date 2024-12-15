package storage

import (
	"context"
	"time"

	"github.com/udevs/lesson3/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrdersStorage struct {
	collection *mongo.Collection
}

func NewOrdersStorage(coll *mongo.Collection) *OrdersStorage {
	return &OrdersStorage{
		collection: coll,
	}
}

func (o *OrdersStorage) Create(ctx context.Context, order *models.Order) (*models.Order, error) {
	order.ID = primitive.NewObjectID().Hex()
	order.CreatedAt = time.Now().Format(time.RFC3339)
	order.UpdatedAt = order.CreatedAt

	_, err := o.collection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *OrdersStorage) FindByID(ctx context.Context, id string) (*models.Order, error) {
	var order models.Order
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = o.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *OrdersStorage) FindAll(ctx context.Context, page, limit int, status string) ([]*models.Order, error) {
	var orders []*models.Order
	query := bson.M{}
	if status != "" {
		query["status"] = status
	}

	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))

	cursor, err := o.collection.Find(ctx, query, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *OrdersStorage) Update(ctx context.Context, id string, order *models.Order) (*models.Order, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	order.UpdatedAt = time.Now().Format(time.RFC3339)

	update := bson.M{"$set": order}
	_, err = o.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *OrdersStorage) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = o.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (o *OrdersStorage) GenerateReport(ctx context.Context, startDate, endDate string) ([]*models.Order, error) {
	var report []*models.Order

	matchStage := bson.D{
		{
			Key: "$match",
			Value: bson.M{
				"created_at": bson.M{
					"$gte": startDate,
					"$lte": endDate,
				},
			},
		},
	}

	groupStage := bson.D{
		{
			Key: "$group",
			Value: bson.M{
				"_id":    "$customer_id",
				"total":  bson.M{"$sum": "$total_price"},
				"orders": bson.M{"$push": "$$ROOT"},
			},
		},
	}

	sortStage := bson.D{
		{
			Key: "$sort",
			Value: bson.M{
				"total": -1,
			},
		},
	}

	cursor, err := o.collection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, sortStage})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result models.Order
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		report = append(report, &result)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return report, nil
}

func (o *OrdersStorage) Count(ctx context.Context, status string) (int64, error) {
	query := bson.M{}
	if status != "" {
		query["status"] = status
	}
	count, err := o.collection.CountDocuments(ctx, query)
	return count, err
}
