package db

import (
	"context"
	"fmt"

	"github.com/Harsh-apk/notesWebApp/types"
	"github.com/Harsh-apk/notesWebApp/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	InsertUser(context.Context, *types.User) error
	GetUser(context.Context, string) (*types.User, error)
	LoginUser(context.Context, *types.IncomingLoginUser) (*types.User, error)
}

type MongoUserStore struct {
	Client *mongo.Client
	Coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		Client: client,
		Coll:   client.Database(DBNAME).Collection(USERCOLL),
	}
}

func (n *MongoUserStore) InsertUser(ctx context.Context, user *types.User) error {
	res, err := n.Coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (n *MongoUserStore) GetUser(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = n.Coll.FindOne(ctx, bson.D{{Key: "_id", Value: oid}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}
func (n *MongoUserStore) LoginUser(ctx context.Context, data *types.IncomingLoginUser) (*types.User, error) {
	var user types.User
	err := n.Coll.FindOne(ctx, bson.D{{Key: "email", Value: data.Email}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	if utils.ComparePassword(&user.EncryptedPassword, &data.Password) {
		return &user, nil
	}
	return nil, fmt.Errorf("invalid email or password")
}
