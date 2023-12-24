package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type DocPayload struct {
	DocHTML string `json:"docHTML"`
}

func GenerateDoc(c *fiber.Ctx) error {
	var payload DocPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	docHTML := payload.DocHTML

	fmt.Printf("Received HTML string from the client: %s\n", docHTML)

	return c.SendString("HTML string received successfully")
}
