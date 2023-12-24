// utils/utils.go

package utils

import (
	"fmt"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GeneratePDF(docHTML string) {
	// PDF_DATA := fmt.Sprintf(`
	// <div style="margin-top:30px">
	// 	<div style="margin-top:10px">%s</div>
	// 	<div style="display: flex; justify-content: center; align-items: center; height: 100vh;">
	// 		<p style="font-size: 12px; color: #888; background-color: white; padding: 0 10px;">
	// 			----------- end of document -----------
	// 		</p>
	// 	</div>
	// </div>
	// `, docHTML)
	// fmt.Printf(docHTML)
	PDF_DATA := "<p>Hello, World!</p>"

	htmlReader := strings.NewReader(PDF_DATA)

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		fmt.Println("Error creating PDF generator:", err)
		return
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(htmlReader))

	err = pdfg.Create()
	if err != nil {
		fmt.Println("Error generating PDF:", err)
		return
	}

	filename := fmt.Sprintf("assets/document.pdf")
	err = pdfg.WriteFile(filename)
	if err != nil {
		fmt.Println("Error writing PDF to file:", err)
		return
	}
}
