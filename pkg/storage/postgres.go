package storage

import (
	"SafeSend/pkg/config"
	"SafeSend/pkg/ent"
	"context"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type Database struct {
	client *ent.Client
}

func (db *Database) Connect() error {

	cfg := config.GetConfig()

	log.Info("Connecting to postgres using host=%s@%s:%d", cfg.Db.User, cfg.Db.Host, cfg.Db.Port)

	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.User,
		cfg.Db.Database,
		cfg.Db.Password,
		cfg.Db.SSLMode,
	)

	client, err := ent.Open("postgres", connectionString)

	if err != nil {
		return err
	}

	defer client.Close()

	// run database migrations
	if err := client.Schema.Create(context.Background()); err != nil {
		return err
	}

	log.Info("Connection to postgres was successfully established")
	db.client = client
	return nil
}

func (db *Database) Close() error {
	return db.client.Close()
}
