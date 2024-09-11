package services

import (
	"file-storage-server/internal/models"

	"gorm.io/gorm"
)

type FileService struct {
	db *gorm.DB
}

func NewFileService(db *gorm.DB) *FileService {
	return &FileService{db: db}
}

func (s *FileService) SaveMetadata(metadata *models.FileMetadata) error {
	return s.db.Create(metadata).Error
}

func (s *FileService) GetAllMetadata() ([]models.FileMetadata, error) {
	var files []models.FileMetadata
	if err := s.db.Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func (s *FileService) GetMetadata(fileID string) (*models.FileMetadata, error) {
	var metadata models.FileMetadata
	if err := s.db.First(&metadata, "id = ?", fileID).Error; err != nil {
		return nil, err
	}
	return &metadata, nil
}
