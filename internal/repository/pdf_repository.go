package repository

import (
	"fmt"
	"strconv"
	"strings"

	"statement_parser/internal/domain"

	"github.com/ledongthuc/pdf"
)

type PDFRepository struct {
	invalidPhrases map[string]struct{}
}

func NewPDFRepository() *PDFRepository {
	return &PDFRepository{
		invalidPhrases: domain.InvalidPhrases,
	}
}

func (r *PDFRepository) CountWordOccurrences(pdfPath, word string) (int, float64, error) {
	f, rdr, err := pdf.Open(pdfPath)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	word = strings.TrimSpace(strings.ToLower(word))
	count := 0
	amount := 0.0
	totalPages := rdr.NumPage()

	for pageIndex := 1; pageIndex <= totalPages; pageIndex++ {
		pageCount, pageAmount, err := r.processPage(rdr, pageIndex, word)
		if err != nil {
			return 0, 0, err
		}
		count += pageCount
		amount += pageAmount
	}

	return count, amount, nil
}

func (r *PDFRepository) processPage(rdr *pdf.Reader, pageIndex int, word string) (int, float64, error) {
	p := rdr.Page(pageIndex)
	if p.V.IsNull() {
		return 0, 0, nil
	}

	content, err := p.GetPlainText(nil)
	if err != nil {
		return 0, 0, err
	}
	count, amount := r.countWordInContent(content, word)
	return count, amount, nil
}

func (r *PDFRepository) countWordInContent(content, word string) (int, float64) {
	lines := strings.Split(content, "\n")
	count := 0
	total := 0.0
	for i := 0; i < len(lines)-1; i++ {
		line := strings.TrimSpace(strings.ToLower(lines[i]))
		if line != word {
			continue
		}

		nextLine := strings.TrimSpace(lines[i+1])
		if _, isInvalid := r.invalidPhrases[nextLine]; !isInvalid {
			amount := lines[i-1]

			cleaned := strings.ReplaceAll(amount, "₸", "")
			cleaned = strings.ReplaceAll(cleaned, " ", "")

			// Replace comma with dot for decimal parsing
			cleaned = strings.ReplaceAll(cleaned, ",", ".")

			// Parse to float64
			value, err := strconv.ParseFloat(cleaned, 64)
			if err != nil {
				fmt.Printf("Error parsing %s: %v\n", amount, err)
				continue
			}
			total += value
			count++
		}
	}

	return count, total
}
