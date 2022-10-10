package database

import (
	"context"
	"reflect"
	"time"

	"buf-connect-playground/utils/otel"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

var client *mongo.Client

// Database is the exported mongo.Database connector
var Database *mongo.Database

// InitializeMongoDB initializes global MongoDB client
func InitializeMongoDB(uri string, db string) (disconnect func() error, err error) {
	timeoutContext, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	structcodec, _ := bsoncodec.NewStructCodec(bsoncodec.JSONFallbackStructTagParser)
	rb := bson.NewRegistryBuilder()
	// register struct codec
	rb.RegisterDefaultEncoder(reflect.Struct, structcodec)
	rb.RegisterDefaultDecoder(reflect.Struct, structcodec)

	client, err = mongo.Connect(timeoutContext,
		options.Client().
			SetRegistry(rb.Build()).
			ApplyURI(uri).
			SetMonitor(otelmongo.NewMonitor(otelmongo.WithTracerProvider(otel.OpenTelTracer))))

	if err != nil {
		client.Disconnect(timeoutContext)
		cancel()
		return nil, err
	}

	if err := client.Ping(timeoutContext, readpref.Primary()); err != nil {
		client.Disconnect(timeoutContext)
		cancel()
		return nil, err
	}

	Database = client.Database(db)

	return func() error {
		err := client.Disconnect(timeoutContext)
		cancel()
		return err
	}, nil
}
