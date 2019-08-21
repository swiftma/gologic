package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateClient -
func CreateClient(uri string) (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return mongo.Connect(ctx, options.Client().ApplyURI(uri))
}