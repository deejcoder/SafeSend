package collections

import (
	storage "SafeSend/storage/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Entity      `bson:",inline"`
	DisplayName string `json:"display_name" bson:"display_name"`
}

func SetCollection(db *storage.Database) *mongo.Collection {
	return db.SetCollection("users")
}

func GetUsers(db *storage.Database) ([]*User, error) {
	col := SetCollection(db)

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
func CreateUser(db *storage.Database, displayName string) (*User, error) {
	col := SetCollection(db)

	user := User{

		DisplayName: displayName,
	}

	result, err := col.InsertOne(db.DefaultContext(), user)
	if err != nil {
		return nil, err
	}

	oid, _ := db.GetObjectId(result)
	user.ID = oid
	return FindUser(db, &user)
}

// FindUser finds a user based on populated attributes
func FindUser(db *storage.Database, user *User) (*User, error) {

	col := SetCollection(db)

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
