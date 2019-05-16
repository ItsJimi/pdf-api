package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/labstack/echo"
)

// PDFOptions define options available to create PDF
type PDFOptions struct {
	Orientation string `json:"orientation" query:"orientation"`
	HTML        string `json:"html" query:"html"`
	URL         string `json:"url" query:"url"`
}

// MSGResponse define response message for users
type MSGResponse struct {
	MSG string `json:"msg"`
}

func generatePDF(options *PDFOptions) (string, *wkhtmltopdf.PDFGenerator) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Print(err)
		return "Can't create pdf", nil
	}

	pdfg.Dpi.Set(300)

	if options.Orientation == "landscape" {
		pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	} else {
		pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	}

	if options.URL != "" {
		pdfg.AddPage(wkhtmltopdf.NewPage(options.URL))
	}
	if options.HTML != "" {
		pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(options.HTML)))
	}

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Print(err)
		return "Can't create pdf", nil
	}

	return "", pdfg
}

func home(c echo.Context) error {
	return c.JSON(http.StatusOK, MSGResponse{
		MSG: "PDF API v0.1.0",
	})
}

func generate(c echo.Context) error {
	options := new(PDFOptions)
	if err := c.Bind(options); err != nil {
		return c.JSON(http.StatusBadRequest, MSGResponse{
			MSG: "Wrong params",
		})
	}
	if options.URL == "" && options.HTML == "" {
		return c.JSON(http.StatusBadRequest, MSGResponse{
			MSG: "Wrong params",
		})
	}
	if options.URL != "" && options.HTML != "" {
		return c.JSON(http.StatusBadRequest, MSGResponse{
			MSG: "Please provide only one source (url or html)",
		})
	}
	err, pdf := generatePDF(options)
	if err != "" {
		return c.JSON(http.StatusInternalServerError, MSGResponse{
			MSG: err,
		})
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/pdf")
	return c.Blob(http.StatusOK, "application/pdf", pdf.Bytes())
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	e := echo.New()
	e.GET("/", home)
	e.GET("/generate", generate)
	e.POST("/generate", generate)
	e.Logger.Fatal(e.Start(":" + port))
}
