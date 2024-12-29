package http

import (
	"context"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"nootebook.com/internal/repository/database"
	"nootebook.com/internal/services"
	"nootebook.com/utils"
	"os"
)

var (
	listenAddr = flag.String("listenAddr", ":5000", "server running properly")
	ctx        = context.Background()
)

func registerRoutes(app *fiber.App) {

	dbHost := os.Getenv("host")
	dbUser := os.Getenv("user")
	dbPass := os.Getenv("pass")
	dbName := os.Getenv("dbname")
	sslmode := os.Getenv("sslmode")
	dbPort := os.Getenv("port")

	fmt.Println(dbHost, dbPort, dbUser, dbPass, dbName, sslmode, 20, 10)
	db, err := utils.PostgresConn(dbHost, dbPort, dbUser, dbPass, dbName, sslmode, 20, 10)
	if err != nil {
		panic(err)
	}
	contactRepo := database.NewContactRepo(db, db)
	contactService := services.NewContactService(contactRepo)
	contactController := ContactController{
		ContactService: contactService,
	}
	v1 := app.Group("/v1")

	v1.Post("/contacts/insert", contactController.Insert)
	v1.Get("/contacts/:name", contactController.Get)
	v1.Get("/contacts", contactController.GetAll)
	v1.Put("/contacts/:name", contactController.Update)
	v1.Delete("/contacts/:name", contactController.Delete)

}
