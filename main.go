package main

import (
	"log"
	"net/http"

	"statement_parser/internal/handler"
	"statement_parser/internal/repository"
	"statement_parser/internal/service"
)

func main() {

	pdfRepo := repository.NewPDFRepository()
	pdfService := service.NewPDFService(pdfRepo)
	countHandler := handler.NewCountHandler(pdfService)

	mux := http.NewServeMux()
	mux.Handle("/count", countHandler)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server listening on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
