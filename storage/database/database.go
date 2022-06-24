package storage

import (
	"SafeSend/config"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Database struct {
	handle *mongo.Database
}

var (
	dbContext Database
)

func (db *Database) Connect() error {

	cfg := config.GetConfig()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	log.Info("Connecting to the database using host=%s:%d", cfg.Db.Host, cfg.Db.Port)

	uri := fmt.Sprintf("mongodb://%s:%d", cfg.Db.Host, cfg.Db.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	// ping the server to ensure we are connected
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	log.Info("Connection to the database was successfully established")
	db.handle = client.Database(cfg.Db.Database)
	return nil
}

func (db *Database) Close() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return db.handle.Client().Disconnect(ctx)
}

func (db *Database) DefaultContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func (db *Database) GetObjectId(insertResult *mongo.InsertOneResult) (primitive.ObjectID, bool) {

	if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		if primitive.IsValidObjectID(oid.String()) {
			return oid, true
		}
	}

	return primitive.NilObjectID, false
}

func Close() error {
	return dbContext.Close()
}

func (db *Database) SetCollection(name string) *mongo.Collection {
	return db.handle.Collection(name)
}
