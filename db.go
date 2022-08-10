package main

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func InsertSlugURLPairToAtlasCollection(slugUrlPair SlugURLPair) *mongo.InsertOneResult {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		panic("mongo_uri cannot be \"\".")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Panic if connection is lost
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Get collection
	coll := client.Database("delongify").Collection("SlugURLPairs")

	// Create index, expires document after SECONDS_TO_EXPIRATION seconds have passed
	index := mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "CreatedAt", Value: bsonx.Int32(1)}},
		Options: options.Index().SetExpireAfterSeconds(SECONDS_TO_EXPIRATION),
	}
	_, err = coll.Indexes().CreateOne(context.TODO(), index)
	if err != nil {
		panic(err)
	}

	// Insert document
	result, err := coll.InsertOne(context.TODO(), slugUrlPair)
	if err != nil {
		panic(err)
	}

	return result
}

func GetSlugURLPair(slug string) (SlugURLPair, error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		panic("mongo_uri cannot be \"\".")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Panic if connection is lost
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Get collection
	coll := client.Database("delongify").Collection("SlugURLPairs")

	var slugURLPair SlugURLPair
	err = coll.FindOne(context.TODO(), bson.D{{Key: "Slug", Value: slug}}).
		Decode(&slugURLPair)

	return slugURLPair, err
}

func SlugIsUnique(slug string) bool {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		panic("mongo_uri cannot be \"\".")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Panic if connection is lost
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Get collection
	coll := client.Database("delongify").Collection("SlugURLPairs")

	var slugURLPair SlugURLPair
	err = coll.FindOne(context.TODO(), bson.D{{Key: "Slug", Value: slug}}).
		Decode(&slugURLPair)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return true
		}
		panic(err)
	}
	return false
}
