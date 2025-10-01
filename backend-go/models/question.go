package models

import "time"

type Question struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PaperID     uint      `gorm:"not null;constraint:OnDelete:CASCADE" json:"paper_id"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	Type        string    `gorm:"type:varchar(20);not null" json:"type"`
	Options     string    `gorm:"type:text" json:"options"`
	Answer      string    `gorm:"type:text;not null" json:"answer"`
	Explanation string    `gorm:"type:text" json:"explanation"`
	ImagePath   string    `gorm:"type:varchar(255)" json:"image_path"`
	CreatedAt   time.Time `json:"created_at"`
	Paper       Paper     `gorm:"foreignKey:PaperID" json:"-"`
}

func (Question) TableName() string {
	return "question"
} 