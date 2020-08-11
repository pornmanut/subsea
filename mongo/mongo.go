package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ErrObjectID is an error. when cannot covert objectID
var ErrObjectID = errors.New("Can not covert object into primitive ObjectID")

// ConnectMongoServer connect to mongo server with url
func ConnectMongoServer(ctx context.Context, url string) (*mongo.Client, error) {

	// url = url + "?retryWrites=false"
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return nil, err
	}

	return client, nil
}
