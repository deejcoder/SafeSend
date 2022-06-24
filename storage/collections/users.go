package collections

import (
	store "SafeSend/storage"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Base
	DisplayName string `json:"display_name" bson:"display_name"`
}

func SetCollection(db *store.Database) *mongo.Collection {
	return db.SetCollection("users")
}

func GetUsers(db *store.Database) ([]*User, error) {
	col := SetCollection(db)

	users := make([]*User, 0)
	cursor, err := col.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return users, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return users, err
		}

		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func (u *User) String() string {
	return u.DisplayName
}
