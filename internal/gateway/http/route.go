package http

import (
	"context"
	"flag"
	"github.com/gofiber/fiber/v2"
	"nootebook.com/config"
	"nootebook.com/internal/repository/database"
	"nootebook.com/internal/services"
	"nootebook.com/utils"
)

var (
	listenAddr = flag.String("listenAddr", ":5000", "server running properly")
	ctx        = context.Background()
)

func registerRoutes(app *fiber.App) {

	db, err := utils.PostgresConn(
		config.AppConfig.Database.Postgres.Host,
		config.AppConfig.Database.Postgres.Port,
		config.AppConfig.Database.Postgres.User,
		config.AppConfig.Database.Postgres.Pass,
		config.AppConfig.Database.Postgres.DBNAME,
		config.AppConfig.Database.Postgres.SSLMODE,
		config.AppConfig.Database.Postgres.MaxOpenConns,
		config.AppConfig.Database.Postgres.MaxIdleConns,
		config.AppConfig.Database.Postgres.Timeout,
	)
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
