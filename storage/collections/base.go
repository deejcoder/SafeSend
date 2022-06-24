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

type Base struct {
	ID           primitive.ObjectID `json:"_id" bson:":"_id"`
	DateCreated  time.Time          `json:"date_created" bson:"date_created"`
	DateModified time.Time          `json:"date_modified" bson:"date_modified"`
}

func (b *Base) CollectionName() string {
	return ""
}

func (b *Base) SetCollection(db *store.Database) *mongo.Collection {
	return db.SetCollection(b.CollectionName())
}
