package services

import (
	"file-storage-server/internal/models"

	"gorm.io/gorm"
)

type ChunkService struct {
	db *gorm.DB
}

func NewChunkService(db *gorm.DB) *ChunkService {
	return &ChunkService{db: db}
}

func (s *ChunkService) SaveChunk(chunk *models.FileChunk) error {
	return s.db.Create(chunk).Error
}

func (s *ChunkService) GetFileChunks(fileID string) ([]*models.FileChunk, error) {
	var chunks []*models.FileChunk
	if err := s.db.Where("file_id = ?", fileID).Order("part").Find(&chunks).Error; err != nil {
		return nil, err
	}
	return chunks, nil
}
