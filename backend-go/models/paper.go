package models

import "time"

type Paper struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Course      string    `gorm:"type:varchar(100);not null" json:"course"`
	Year        int       `gorm:"not null" json:"year"`
	Term        string    `gorm:"type:varchar(20)" json:"term"`
	College     string    `gorm:"type:varchar(100)" json:"college"`
	ImagePath   string    `gorm:"type:varchar(255)" json:"image_path"`
	FileType    string    `gorm:"type:varchar(20);default:image" json:"file_type"`
	FileSize    int       `gorm:"default:0" json:"file_size"`
	StorageType string    `gorm:"type:varchar(20);default:local" json:"storage_type"`
	CreatedAt   time.Time `json:"created_at"`
}

func (Paper) TableName() string {
	return "paper"
} 