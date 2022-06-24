package collections

import (
	store "SafeSend/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Collection interface {
	CollectionName()
	SetCollection(driver *store.Database)
}

type Entity struct {
	ID           primitive.ObjectID `json:"id" bson:":"_id"`
	DateCreated  time.Time          `json:"date_created" bson:"date_created"`
	DateModified time.Time          `json:"date_modified" bson:"date_modified"`
}

func (b *Entity) CollectionName() string {
	return ""
}

func (b *Entity) SetCollection(db *store.Database) *mongo.Collection {
	return db.SetCollection(b.CollectionName())
}

func CreateEntity() *Entity {
	return &Entity{
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
}
