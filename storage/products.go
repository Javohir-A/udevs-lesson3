/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-12-14 04:17:57
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-15 05:57:37
 * @FilePath: /lesson3/storage/products.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package storage

import (
	"context"
	"log"
	"time"

	"github.com/udevs/lesson3/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductStorage struct {
	collection *mongo.Collection
}

func NewProductStorage(coll *mongo.Collection) *ProductStorage {
	return &ProductStorage{
		collection: coll,
	}
}

func (p *ProductStorage) Create(ctx context.Context, product *models.Product) (*models.Product, error) {
	curTime := time.Now().Format("2006-01-02")

	res, err := p.collection.InsertOne(ctx, bson.D{
		{Key: "name", Value: product.Name},
		{Key: "price", Value: product.Price},
		{Key: "stock", Value: product.Stock},
		{Key: "category", Value: product.Category},
		{Key: "created_at", Value: curTime},
	})

	if err != nil {
		return nil, err
	}

	objID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Println("ObjectId is not valid!")
		return nil, err
	}

	return &models.Product{
		ID:        objID.Hex(),
		Name:      product.Name,
		Category:  product.Category,
		Stock:     product.Stock,
		Price:     product.Price,
		CreatedAt: curTime,
	}, nil
}

func (p *ProductStorage) FindByID(ctx context.Context, id string) (*models.Product, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	res := p.collection.FindOne(ctx, bson.D{
		{Key: "_id", Value: objID},
	})
	prod := models.Product{}

	if err := res.Decode(&prod); err != nil {
		return nil, err
	}

	return &prod, nil
}

func (p *ProductStorage) FindAll(ctx context.Context, page, limit int, search string) ([]*models.Product, error) {
	var products []*models.Product

	filter := bson.D{}
	if search != "" {
		filter = bson.D{
			{Key: "name", Value: bson.D{{Key: "$regex", Value: search}, {Key: "$options", Value: "i"}}},
		}
	}

	opts := &options.FindOptions{}
	opts.SetSkip(int64((page - 1) * limit))
	opts.SetLimit(int64(limit))

	cursor, err := p.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func (p *ProductStorage) Update(ctx context.Context, id string, product *models.Product) (*models.Product, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: product.Name},
			{Key: "price", Value: product.Price},
			{Key: "stock", Value: product.Stock},
			{Key: "category", Value: product.Category},
			{Key: "updated_at", Value: primitive.DateTime(time.Now().Unix())},
		}},
	}

	_, err = p.collection.UpdateByID(ctx, objID, update)
	if err != nil {
		return nil, err
	}

	return p.FindByID(ctx, id)
}

func (p *ProductStorage) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = p.collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: objID}})
	return err
}

func (p *ProductStorage) Count(ctx context.Context, search string) (int64, error) {
	filter := bson.D{}
	if search != "" {
		filter = bson.D{
			{Key: "name", Value: bson.D{{Key: "$regex", Value: search}, {Key: "$options", Value: "i"}}},
		}
	}
	count, err := p.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}
