package service

import (
	"os"

	"statement_parser/internal/domain"
)

type PDFService struct {
	repo domain.PDFProcessor
}

func NewPDFService(repo domain.PDFProcessor) *PDFService {
	return &PDFService{repo: repo}
}

func (s *PDFService) CountWordFromFile(filePath, word string) (int, error) {
	return s.repo.CountWordOccurrences(filePath, word)
}

func (s *PDFService) SaveUploadedFile(fileContent []byte) (string, error) {
	tmpFile, err := os.CreateTemp("", "*.pdf")
	if err != nil {
		return "", err
	}

	if _, err := tmpFile.Write(fileContent); err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return "", err
	}

	tmpFile.Close()
	return tmpFile.Name(), nil
}

func (s *PDFService) CleanupFile(filePath string) {
	os.Remove(filePath)
}
