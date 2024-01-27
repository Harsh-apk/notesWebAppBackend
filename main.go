package main

import (
	"context"
	"flag"
	"log"

	"github.com/Harsh-apk/notesWebApp/api"
	"github.com/Harsh-apk/notesWebApp/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUri = "mongodb://localhost:27017"

var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", "localhost:5000", "The Listen Address of the api server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	notesHandler := api.NewNotesHandler(db.NewMongoNotesStore(client))

	app := fiber.New(config)
	app.Use(cors.New())
	apiv1 := app.Group("/api/v1")
	//user methods
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Post("/user/login", userHandler.HandleLoginUser)
	apiv1.Get("/user/:id/user", userHandler.HandleGetUser)
	//notes methods
	apiv1.Post("/user/:id", notesHandler.HandlePostNotes)
	apiv1.Get("/user/:id", notesHandler.HandleGetNotes)
	apiv1.Delete("/user/:noteId", notesHandler.HandleDeleteNote)
	app.Static("/", "./public/build")
	app.Listen(*listenAddr)

}
