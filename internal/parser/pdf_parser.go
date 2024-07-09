package parser

import (
	"fmt"
	"github.com/unidoc/unipdf/v3/contentstream"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
	"log"
)

func ParsePDF(filePath string) (string, error) {
	// Read the PDF file
	pdfReader, _, err := model.NewPdfReaderFromFile(filePath, nil)
	if err != nil {
		return "", fmt.Errorf("could not read PDF file: %v", err)
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		log.Fatalf("Failed to get number of pages: %v\n", err)
	}

	fmt.Printf("No. of pages collected from PDF: %d\n", numPages)

	var text string

	for i := 1; i <= numPages; i++ {
		page, err := pdfReader.GetPage(i)
		if err != nil {
			return "", fmt.Errorf("could not get page '%d': %v\n", i, err)
		}
		contentStreams, err := page.GetContentStreams()
		if err != nil {
			log.Printf("Error extracting text from page '%d': %v\n", i, err)
		}

		pageContentStr := ""

		for _, contentStream := range contentStreams {
			pageContentStr += contentStream
		}

		contentStreamParser := contentstream.NewContentStreamParser(pageContentStr)

		operations, err := contentStreamParser.Parse()
		if err != nil {
			log.Printf("Error parsing content stream: %v\n", err)
		}

		for i, op := range *operations {
			log.Printf("Operation %d: %s - Params: %v\n", i+1, op.Operand, op.Params)
		}

		ex, err := extractor.New(page)
		if err != nil {
			log.Printf("Instantiation of extractor failed: %v\n", err)
		}

		text, err = ex.ExtractText()
		log.Printf("Text: %v", text)

	}

	return text, nil
}
