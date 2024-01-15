package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"practice/forms"
	"strings"
)

func Convert(c *fiber.Ctx) error {

	wkhtmltopdf.SetPath("C:\\Program Files\\wkhtmltopdf\\bin\\wkhtmltopdf.exe")

	var requestBody forms.HTMLRequestBody
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Geçersiz istek formatı")
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Dönüştürücü başlatılamadı: %v", err))
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(requestBody.HTMLContent)))

	err = pdfg.Create()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("PDF oluşturulamadı: %v", err))
	}

	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "inline; filename=output.pdf")

	return c.Status(fiber.StatusOK).Send(pdfg.Bytes())
}
