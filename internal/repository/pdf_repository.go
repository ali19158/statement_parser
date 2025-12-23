package repository

import (
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

func (r *PDFRepository) CountWordOccurrences(pdfPath, word string) (int, error) {
	f, rdr, err := pdf.Open(pdfPath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	word = strings.TrimSpace(strings.ToLower(word))
	count := 0
	totalPages := rdr.NumPage()

	for pageIndex := 1; pageIndex <= totalPages; pageIndex++ {
		pageCount, err := r.processPage(rdr, pageIndex, word)
		if err != nil {
			return 0, err
		}
		count += pageCount
	}

	return count, nil
}

func (r *PDFRepository) processPage(rdr *pdf.Reader, pageIndex int, word string) (int, error) {
	p := rdr.Page(pageIndex)
	if p.V.IsNull() {
		return 0, nil
	}

	content, err := p.GetPlainText(nil)
	if err != nil {
		return 0, err
	}

	return r.countWordInContent(content, word), nil
}

func (r *PDFRepository) countWordInContent(content, word string) int {
	lines := strings.Split(content, "\n")
	count := 0

	for i := 0; i < len(lines)-1; i++ {
		line := strings.TrimSpace(strings.ToLower(lines[i]))
		if line != word {
			continue
		}

		nextLine := strings.TrimSpace(lines[i+1])
		if _, isInvalid := r.invalidPhrases[nextLine]; !isInvalid {
			count++
		}
	}

	return count
}
