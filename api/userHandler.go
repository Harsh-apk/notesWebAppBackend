package api

import (
	"github.com/Harsh-apk/notesWebApp/db"
	"github.com/Harsh-apk/notesWebApp/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		UserStore: userStore,
	}
}

func (u *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var data *types.IncomingUser
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	error := data.ValidateUser()
	if len(*error) != 0 {
		return c.JSON(error)
	}
	user, err := types.UserFromIncomingUser(data)
	if err != nil {
		return err
	}
	err = u.UserStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
func (u *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := u.UserStore.GetUser(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(res)
}
func (u *UserHandler) HandleLoginUser(c *fiber.Ctx) error {
	var data *types.IncomingLoginUser
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	user, err := u.UserStore.LoginUser(c.Context(), data)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
