package api

import (
	"encoding/json"
	"file-storage-server/internal/models"
	"file-storage-server/internal/services"
	"file-storage-server/pkg/utils"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Handlers struct {
	fileService  *services.FileService
	chunkService *services.ChunkService
}

func NewHandlers(fs *services.FileService, cs *services.ChunkService) *Handlers {
	return &Handlers{
		fileService:  fs,
		chunkService: cs,
	}
}

func (h *Handlers) UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileID := utils.GenerateUUID()
	metadata := &models.FileMetadata{
		ID:   fileID,
		Name: header.Filename,
		Size: header.Size,
	}

	if err := h.fileService.SaveMetadata(metadata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	chunkSize := 1024 * 1024 // 1MB chunks
	partNum := 0

	for {
		chunk := make([]byte, chunkSize)
		n, err := file.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fileChunk := &models.FileChunk{
			FileID: fileID,
			Data:   chunk[:n],
			Part:   partNum,
		}

		if err := h.chunkService.SaveChunk(fileChunk); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		partNum++
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"file_id": fileID})
}

func (h *Handlers) GetFilesHandler(w http.ResponseWriter, r *http.Request) {
	files, err := h.fileService.GetAllMetadata()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(files)
}

func (h *Handlers) DownloadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileID := vars["id"]

	metadata, err := h.fileService.GetMetadata(fileID)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	chunks, err := h.chunkService.GetFileChunks(fileID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", metadata.Name))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", metadata.Size))

	for _, chunk := range chunks {
		if _, err := w.Write(chunk.Data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
