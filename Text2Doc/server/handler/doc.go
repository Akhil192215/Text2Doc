package handler

import (
	"fmt"
	"log"
	"os"
	"text2doc/utils"

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
	utils.GeneratePDF(docHTML)

	pdfFilePath := fmt.Sprintf("assets/document.pdf")
	pdfData, err := os.ReadFile(pdfFilePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error reading PDF file")
	}
	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename=meeting.pdf")
	defer func() {
		if err := os.Remove(pdfFilePath); err != nil {
			log.Println("Error deleting PDF file:", err)
		}
	}()

	return c.Send(pdfData)
}
