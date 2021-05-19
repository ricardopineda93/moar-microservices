package mongodb

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/rjjp5294/url-shortener/shortener"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoUrl string, mongoTimeout int) (*mongo.Client, error) {
	// Set up the context to use when connecting with MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	// This ensures proper timeout from out context
	defer cancel()
	// Create a MongoDB client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return nil, err
	}
	// Ping the Primary cluster to ensure the connection is established
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	// Return the client and
	return client, nil
}

// Factory function for creating a MongoDB repository instance that implements the
// shortener/repository interface... Just handles the setup of the MongoDB connection and client
// and returns the fully implemented interface
func NewMongoRepository(mongoUrl, mongoDB string, mongoTimeout int) (shortener.RedirectRepository, error) {
	repo := &mongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}
	client, err := newMongoClient(mongoUrl, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepository")
	}
	repo.client = client

	return repo, nil
}

func (r *mongoRepository) Find(code string) (*shortener.Redirect, error) {
	// Setting up the context to be used for the MongoDB query
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	// Initialize pointer to a shortener struct instance to be "casted" to from MongoDB result
	redirect := &shortener.Redirect{}
	// Get the collection we want to query from our DB
	collection := r.client.Database(r.database).Collection("redirects")
	// Declare our query filter as a BSON document
	filter := bson.M{"code": code}
	// Attempt to find in our collection, cast to our redirect struct
	err := collection.FindOne(ctx, filter).Decode(&redirect)
	// Handle errors is any arise from our MongoDB query
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(shortener.ErrRedirectNotFound, "redirects")
		}
		return nil, errors.Wrap(err, "repository.Redirect.Find")
	}
	// Return
	return redirect, nil
}

func (r *mongoRepository) Store(redirect *shortener.Redirect) error {
	// Setting up the context to be used for the MongoDB query
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("redirects")
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"code":       redirect.Code,
			"url":        redirect.URL,
			"created_at": redirect.CreatedAt,
		},
	)
	if err != nil {
		return errors.Wrap(err, "repository.Redirect.Store")
	}
	return nil
}
