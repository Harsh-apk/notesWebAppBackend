package db

import (
	"context"
	"fmt"

	"github.com/Harsh-apk/notesWebApp/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotesStore interface {
	InsertNotes(context.Context, *types.Note) error
	GetNotes(context.Context, *primitive.ObjectID) (*[]types.Note, error)
	DeleteNote(context.Context, *primitive.ObjectID) error
}

type MongoNotesStore struct {
	Client *mongo.Client
	Coll   *mongo.Collection
}

func NewMongoNotesStore(client *mongo.Client) *MongoNotesStore {
	return &MongoNotesStore{
		Client: client,
		Coll:   client.Database(DBNAME).Collection(NOTESCOLL),
	}
}

func (n *MongoNotesStore) InsertNotes(ctx context.Context, data *types.Note) error {
	note, err := n.Coll.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	data.NoteID = note.InsertedID.(primitive.ObjectID)
	return nil
}

func (n *MongoNotesStore) GetNotes(ctx context.Context, id *primitive.ObjectID) (*[]types.Note, error) {
	cur, err := n.Coll.Find(ctx, bson.D{{Key: "userId", Value: id}})
	if err != nil {
		return nil, err
	}
	var notes []types.Note
	err = cur.All(ctx, &notes)
	if err != nil {
		return nil, err
	}
	return &notes, nil
}
func (n *MongoNotesStore) DeleteNote(ctx context.Context, id *primitive.ObjectID) error {
	res, err := n.Coll.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return err
	}
	if res.DeletedCount > 0 {
		return nil
	}
	return fmt.Errorf("something went wrong")
}
