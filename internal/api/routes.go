package api

import (
	"file-storage-server/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *mux.Router {
	fileService := services.NewFileService(db)
	chunkService := services.NewChunkService(db)
	handlers := NewHandlers(fileService, chunkService)

	r := mux.NewRouter()
	r.HandleFunc("/upload", handlers.UploadHandler).Methods("POST")
	r.HandleFunc("/files", handlers.GetFilesHandler).Methods("GET")
	r.HandleFunc("/download/{id}", handlers.DownloadHandler).Methods("GET")

	return r
}
