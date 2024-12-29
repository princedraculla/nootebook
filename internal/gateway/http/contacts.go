package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"nootebook.com/internal/services"
	service_models "nootebook.com/internal/services/service_models/types"
)

type ContactController struct {
	ContactService services.Contact
}

type createContactParams struct {
	PhoneNumbers []service_models.PhoneNumber `json:"phone_numbers"`
	Name         string                       `json:"name"`
}
type UpdateContactParams struct {
	phoneNumber []service_models.PhoneNumber
}

func NewContactController(contactService services.Contact) *ContactController {
	return &ContactController{
		ContactService: contactService,
	}
}
func (c *ContactController) Insert(ctx *fiber.Ctx) error {
	var params createContactParams
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	fmt.Printf("body is %+v\n", params)

	err := c.ContactService.Insert(ctx.Context(), &service_models.Contact{
		PhoneNumbers: params.PhoneNumbers,
		Name:         params.Name,
	})
	if err != nil {
		return err
	}
	return ctx.JSON(err)
}

func (c *ContactController) Get(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	contact, err := c.ContactService.Get(ctx.Context(), name)
	if err != nil {
		return err
	}
	return ctx.JSON(contact)
}

func (c *ContactController) GetAll(ctx *fiber.Ctx) error {
	contacts, err := c.ContactService.GetAll(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(contacts)
}

func (c *ContactController) Delete(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	err := c.ContactService.Delete(ctx.Context(), &name)
	if err != nil {
		return err
	}
	msg := map[string]string{
		"status": "success",
	}

	return ctx.JSON(msg)
}

func (c *ContactController) Update(ctx *fiber.Ctx) error {
	var updateParams UpdateContactParams
	newPhoneNumber := ctx.BodyParser(&updateParams)
	updatedPhoneNumber, err := c.ContactService.Update(ctx.Context(), newPhoneNumber)
	if err != nil {
		return err
	}
	return ctx.JSON(updatedPhoneNumber)
}
