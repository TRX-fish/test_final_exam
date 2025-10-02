package services

import (
	"backend-go/database"
	"backend-go/models"
)

type PaperInput struct {
	Course    string `json:"course"`
	Year      int    `json:"year"`
	Term      string `json:"term"`
	College   string `json:"college"`
	ImagePath string `json:"image_path"`
	FileType  string `json:"file_type"`
	FileSize  int    `json:"file_size"`
	FileURL   string `json:"file_url"`
}

func AddPaper(input PaperInput) (*models.Paper, error) {
	paper := models.Paper{
		Course:    input.Course,
		Year:      input.Year,
		Term:      input.Term,
		College:   input.College,
		ImagePath: input.ImagePath,
	}

	if err := database.DB.Create(&paper).Error; err != nil {
		return nil, err
	}

	return &paper, nil
}

func UpdatePaperImage(paperID uint, imagePath string) (*models.Paper, error) {
	var paper models.Paper
	if err := database.DB.First(&paper, paperID).Error; err != nil {
		return nil, err
	}

	paper.ImagePath = imagePath
	if err := database.DB.Save(&paper).Error; err != nil {
		return nil, err
	}

	return &paper, nil
}

func GetPapers(course, term, college, keyword string, year, page, perPage int) ([]models.Paper, int64, error) {
	var papers []models.Paper
	var total int64

	query := database.DB.Model(&models.Paper{})

	if course != "" {
		query = query.Where("course LIKE ?", "%"+course+"%")
	}
	if year > 0 {
		query = query.Where("year = ?", year)
	}
	if term != "" {
		query = query.Where("term LIKE ?", "%"+term+"%")
	}
	if college != "" {
		query = query.Where("college LIKE ?", "%"+college+"%")
	}
	if keyword != "" {
		query = query.Where("course LIKE ? OR term LIKE ? OR college LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	query.Count(&total)

	offset := (page - 1) * perPage
	if err := query.Order("created_at DESC").Offset(offset).Limit(perPage).Find(&papers).Error; err != nil {
		return nil, 0, err
	}

	return papers, total, nil
} 