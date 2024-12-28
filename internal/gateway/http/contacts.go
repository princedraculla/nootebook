package http

import (
	"github.com/gofiber/fiber/v2"
	"nootebook.com/internal/services"
	service_models "nootebook.com/internal/services/service_models/types"
)

type Contacts struct {
	ContactService services.Contact
}

type createContactParams struct {
	PhoneNunmber []service_models.PhoneNumber
	Name         string
}

func NewContacts(contactService services.Contact) *Contacts {
	return &Contacts{
		ContactService: contactService,
	}
}
func (c *Contacts) Insert(ctx *fiber.Ctx) error {
	var params createContactParams
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	err := c.ContactService.Insert(ctx.Context(), &service_models.Contact{
		PhoneNumbers: params.PhoneNunmber,
		Name:         params.Name,
	})
	if err != nil {
		return err
	}
	return ctx.JSON(err)
}

func (c *Contacts) Get(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	contact, err := c.ContactService.Get(ctx.Context(), name)
	if err != nil {
		return err
	}
	return ctx.JSON(contact)
}

func (c *Contacts) GetAll(ctx *fiber.Ctx) error {
	contacts, err := c.ContactService.GetAll(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(contacts)
}

func (c *Contacts) 