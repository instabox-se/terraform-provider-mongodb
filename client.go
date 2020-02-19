package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	client *mongo.Client
	ctx    *context.Context
}

type CreateResult struct {
	Ok int `bson:"ok"`
}

func (c *CreateResult) IsOk() bool {
	return c.Ok == 1
}

func NewClient(connectionString string) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Printf("[ERROR] Error creating mongo client %s", err)
		return nil, err
	}

	client := &Client{
		client: c,
		ctx:    &ctx,
	}

	return client, nil
}
