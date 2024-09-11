package models

import "gorm.io/gorm"

type FileChunk struct {
	gorm.Model
	FileID string `gorm:"index"`
	Data   []byte
	Part   int `gorm:"index"`
}

type FileMetadata struct {
	ID   string `gorm:"primaryKey"`
	Name string
	Size int64
}
