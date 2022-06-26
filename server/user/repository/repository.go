package repository

import (
	"SafeSend/server/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type User struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	DisplayName  string             `json:"display_name" bson:"display_name"`
	LastAccessed time.Time          `json:"last_accessed" bson:"last_accessed"`
	DateCreated  time.Time          `json:"date_created" bson:"date_created"`
	DateModified time.Time          `json:"date_modified" bson:"date_modified"`
}

const collectionName = "user"

func setCollection(db *storage.Database) *mongo.Collection {
	return db.SetCollection(collectionName)
}

func GetUsers(db *storage.Database) ([]*User, error) {
	col := setCollection(db)

	users := make([]*User, 0)
	cursor, err := col.Find(db.DefaultContext(), bson.D{{}})
	if err != nil {
		return users, err
	}

	defer cursor.Close(db.DefaultContext())

	for cursor.Next(db.DefaultContext()) {
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

// CreateUser creates a new user and returns it
func CreateUser(db *storage.Database, displayName string) error {
	col := setCollection(db)

	user := User{
		ID:           primitive.NewObjectID(),
		DisplayName:  displayName,
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}

	_, err := col.InsertOne(db.DefaultContext(), user)
	if err != nil {
		return err
	}

	return nil
}

// FindUser finds a user based on populated attributes
func FindUser(db *storage.Database, user *User) (*User, error) {
	col := setCollection(db)

	result := col.FindOne(db.DefaultContext(), user)
	if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) String() string {
	return u.DisplayName
}
