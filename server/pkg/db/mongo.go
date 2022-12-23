package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	c   *mongo.Client
	db  *mongo.Database
	ctx context.Context
}

func NewMongo(uri string, dbName string) (*MongoDB, error) {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	db := client.Database(dbName)
	return &MongoDB{
		c:   client,
		db:  db,
		ctx: ctx,
	}, nil
}

func (m *MongoDB) GetCollection(n string) *mongo.Collection {
	return m.db.Collection(n)
}

func (m *MongoDB) TransformId(id string) primitive.ObjectID {
	i, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	return i
}