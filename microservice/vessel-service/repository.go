package main

import (
	"context"
	"fmt"
	pb "github.com/gologic/microservice/vessel-service/proto/vessel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type repository interface {
	FindAvailable(spec *pb.Specification) (*pb.Vessel, error)
	Create(vessel *pb.Vessel) error
}

type VesselRepository struct {
	collection *mongo.Collection
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repository *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	filter := bson.D{
		{"capacity", bson.D{{"$gte", spec.Capacity}}},
		{"maxweight", bson.D{{"$gte", spec.MaxWeight}}},
	}

	log.Printf("find available spec: %v\n", spec)

	var vessel *pb.Vessel
	if err := repository.collection.FindOne(context.TODO(), filter).Decode(&vessel); err != nil {
		log.Printf("FindOne %v\n", err)

		cur, err := repository.collection.Find(context.TODO(), bson.D{})
		for cur != nil && cur.Next(context.TODO()) {
			err = cur.Decode(&vessel)
			if err == nil {
				return vessel, err
			}
		}
		return nil, fmt.Errorf("no documents at all")
	}
	return vessel, nil
}

// Create a new vessel
func (repository *VesselRepository) Create(vessel *pb.Vessel) error {
	_, err := repository.collection.InsertOne(context.TODO(), vessel)
	return err
}
