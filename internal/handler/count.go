package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"statement_parser/internal/domain"
	"statement_parser/internal/service"
)

type CountHandler struct {
	pdfService *service.PDFService
}

func NewCountHandler(pdfService *service.PDFService) *CountHandler {
	return &CountHandler{pdfService: pdfService}
}

func (h *CountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handleCount(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CountHandler) handleCount(w http.ResponseWriter, r *http.Request) {

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		h.respondWithError(w, "missing file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if !isValidPDF(fileHeader) {
		h.respondWithError(w, "invalid file type", http.StatusBadRequest)
		return
	}

	fileContent, err := io.ReadAll(file)
	if err != nil {
		h.respondWithError(w, "failed to read file", http.StatusInternalServerError)
		return
	}

	tmpFilePath, err := h.pdfService.SaveUploadedFile(fileContent)
	if err != nil {
		h.respondWithError(w, "failed to process file", http.StatusInternalServerError)
		return
	}
	defer h.pdfService.CleanupFile(tmpFilePath)

	count, amount, err := h.pdfService.CountWordFromFile(tmpFilePath)
	if err != nil {
		h.respondWithError(w, "failed to process PDF", http.StatusInternalServerError)
		return
	}

	h.respondWithJSON(w, http.StatusOK, domain.CountResponse{
		Count:  count,
		Amount: fmt.Sprintf("%.2f", amount),
	})
}

func (h *CountHandler) respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func (h *CountHandler) respondWithError(w http.ResponseWriter, message string, status int) {
	h.respondWithJSON(w, status, domain.CountResponse{
		Error: message,
	})
}

func isValidPDF(fileHeader *multipart.FileHeader) bool {
	return strings.HasSuffix(strings.ToLower(fileHeader.Filename), ".pdf")
}
