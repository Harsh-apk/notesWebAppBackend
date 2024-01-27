package api

import (
	"github.com/Harsh-apk/notesWebApp/db"
	"github.com/Harsh-apk/notesWebApp/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotesHandler struct {
	NotesStore db.NotesStore
}

func NewNotesHandler(notesStore db.NotesStore) *NotesHandler {
	return &NotesHandler{
		NotesStore: notesStore,
	}
}

func (n *NotesHandler) HandlePostNotes(c *fiber.Ctx) error {
	var data types.IncomingNote
	data.UserID = c.Params("id")

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	note, err := types.NoteFromIncomingNote(&data)
	if err != nil {
		return err
	}
	err = n.NotesStore.InsertNotes(c.Context(), note)
	if err != nil {
		return nil
	}
	return c.JSON(note)

}
func (n *NotesHandler) HandleGetNotes(c *fiber.Ctx) error {
	uid, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return err
	}
	notes, err := n.NotesStore.GetNotes(c.Context(), &uid)
	if err != nil {
		return err
	}
	if len(*notes) > 0 {
		return c.JSON(notes)
	}
	return c.JSON(map[string]string{"result": "no data found"})

}
func (n *NotesHandler) HandleDeleteNote(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("noteId"))
	if err != nil {
		return err
	}
	err = n.NotesStore.DeleteNote(c.Context(), &id)
	if err != nil {
		return err
	}
	return c.JSON(map[string]string{"result": "success"})

}
