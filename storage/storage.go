package storage

import (
	"SafeSend/config"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	handle *mongo.Database
}

var (
	dbContext Database
)

func (db *Database) Connect() error {

	cfg := config.GetConfig()

	log.Info("Connecting to the database using host=%s:%d", cfg.Db.Host, cfg.Db.Port)

	uri := fmt.Sprintf("mongodb://%s:%d", cfg.Db.Host, cfg.Db.Port)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	// ping the server to ensure we are connected
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return err
	}

	log.Info("Connection to the database was successfully established")
	db.handle = client.Database(cfg.Db.Database)
	return nil
}

func (db *Database) Close() error {
	return db.handle.Client().Disconnect(context.TODO())
}

func Close() error {
	return dbContext.Close()
}

func (db *Database) SetCollection(name string) *mongo.Collection {
	return db.handle.Collection(name)
}
