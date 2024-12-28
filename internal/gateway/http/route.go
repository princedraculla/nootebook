package http

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"nootebook.com/internal/services"
)

var (
	ContactRepo    = "any"
	contactService = services.NewContactService(&ContactRepo)
	ctx            = context.Background()
)
var
func registerRoutes(app *fiber.App) {
	contactController := "sa"
	apiV1 := app.Group("/v1")

	apiV1.Get("/contacts", )
}
