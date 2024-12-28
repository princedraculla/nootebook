package http

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"nootebook.com/internal/services"
)

var (
	ContactRepo       = "any"
	contactService    = services.NewContactService(&ContactRepo)
	ctx               = context.Background()
	contactController = NewContactController(contactService)
)

func registerRoutes(app *fiber.App) {
	v1 := app.Group("/v1")

	v1.Post("/insert", contactController.Insert)
	v1.Get("/contact/:name", contactController.Get)
	v1.Get("/contacts", contactController.GetAll)
	v1.Put("/contact/:name", contactController.Update)
	v1.Delete("/contact/:name", contactController.Delete)

}
