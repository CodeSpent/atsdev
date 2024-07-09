package parser

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func ParsePDF(filePath string) (string, error) {
	// Check if pdfinfo is installed
	if _, err := exec.LookPath("pdfinfo"); err != nil {
		return "", fmt.Errorf("pdfinfo not found in $PATH: %v", err)
	}

	// Check if magick (ImageMagick) is installed
	if _, err := exec.LookPath("magick"); err != nil {
		return "", fmt.Errorf("ImageMagick not found in $PATH: %v", err)
	}

	client := gosseract.NewClient()
	defer client.Close()

	client.SetLanguage("eng")

	pdfPath, err := filepath.Abs(filePath)
	if err != nil {
		return "", fmt.Errorf("could not get absolute path: %v", err)
	}

	tempDir := filepath.Join(os.TempDir(), "pdf_images")
	err = os.MkdirAll(tempDir, 0755)
	if err != nil {
		return "", fmt.Errorf("could not create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	var pdfImages []string
	pageCount, err := getPageCount(pdfPath)
	if err != nil {
		return "", fmt.Errorf("could not get PDF page count: %v", err)
	}

	for i := 0; i < pageCount; i++ {
		imagePath := filepath.Join(tempDir, fmt.Sprintf("page-%d.png", i))
		/*
			FIXME: Magick already supports multiple page PDFs.
			Iterating is creating duplicate files.
		*/
		err := convertPDFToImage(pdfPath, filepath.Join(tempDir, "page.png"))
		if err != nil {
			return "", fmt.Errorf("could not convert PDF page %d to image: %v", i, err)
		}
		pdfImages = append(pdfImages, imagePath)
		fmt.Printf("Converted PDF page %d to image: %s\n", i, imagePath)
	}

	var textBuilder strings.Builder
	for _, imagePath := range pdfImages {
		fmt.Printf("Processing image: %s\n", imagePath)

		if _, err := os.Stat(imagePath); os.IsNotExist(err) {
			return "", fmt.Errorf("image file does not exist: %s", imagePath)
		}

		client.SetImage(imagePath)
		text, err := client.Text()
		if err != nil {
			return "", fmt.Errorf("could not extract text from image: %v", err)
		}
		fmt.Printf("Extracted text: %s\n", text)
		textBuilder.WriteString(text)
		textBuilder.WriteString("\n")
	}

	return textBuilder.String(), nil
}

func getPageCount(pdfPath string) (int, error) {
	cmd := exec.Command("pdfinfo", pdfPath)
	out, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("error running pdfinfo: %v", err)
	}

	info := string(out)
	lines := strings.Split(info, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Pages:") {
			pageCountStr := strings.TrimSpace(strings.TrimPrefix(line, "Pages:"))
			pageCount, err := strconv.Atoi(pageCountStr)
			if err != nil {
				return 0, fmt.Errorf("could not parse page count: %v", err)
			}
			return pageCount, nil
		}
	}

	return 0, fmt.Errorf("could not find page count")
}

func convertPDFToImage(pdfPath, imagePath string) error {
	cmd := exec.Command("convert", "-density", "300", pdfPath, imagePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error converting PDF to image: %v, output: %s", err, string(output))
	}
	return nil
}
