package groups

import (
	"SafeSend/storage/collections/users"
	storage "SafeSend/storage/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Group struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	GroupName       string             `json:"group_name" bson:"group_name"`
	MaxParticipants int                `json:"max_participants" bson:"max_participants"`
	InviteOnly      bool               `json:"invite_only" bson:"invite_only"`
	Users           []users.User
	DateCreated     time.Time `json:"date_created" bson:"date_created"`
	DateModified    time.Time `json:"date_modified" bson:"date_modified"`
}

func setCollection(db *storage.Database) *mongo.Collection {
	return db.SetCollection("groups")
}

func GetGroups(db *storage.Database) ([]*Group, error) {
	col := setCollection(db)

	groups := make([]*Group, 0)
	cursor, err := col.Find(db.DefaultContext(), bson.D{{}})
	if err != nil {
		return groups, err
	}

	defer cursor.Close(db.DefaultContext())

	for cursor.Next(db.DefaultContext()) {
		var group Group
		if err := cursor.Decode(&group); err != nil {
			return groups, err
		}

		groups = append(groups, &group)
	}

	if err := cursor.Err(); err != nil {
		return groups, err
	}

	return groups, nil
}

// CreateGroup creates a new group
func CreateGroup(db *storage.Database, groupName string, maxParticipants int, inviteOnly bool) error {
	col := setCollection(db)

	group := Group{
		ID:              primitive.NewObjectID(),
		GroupName:       groupName,
		MaxParticipants: maxParticipants,
		InviteOnly:      inviteOnly,
		DateCreated:     time.Now(),
		DateModified:    time.Now(),
	}

	_, err := col.InsertOne(db.DefaultContext(), group)
	if err != nil {
		return err
	}

	return nil
}

// FindGroup finds a group based on populated attributes
func FindGroup(db *storage.Database, group *Group) (*Group, error) {
	col := setCollection(db)

	result := col.FindOne(db.DefaultContext(), group)
	if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Decode(group); err != nil {
		return nil, err
	}

	return group, nil
}

func (g *Group) String() string {
	return g.GroupName
}
