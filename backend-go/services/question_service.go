package services

import (
	"backend-go/database"
	"backend-go/models"
)

type QuestionInput struct {
	PaperID     uint   `json:"paper_id"`
	Content     string `json:"content"`
	Type        string `json:"type"`
	Options     string `json:"options"`
	Answer      string `json:"answer"`
	Explanation string `json:"explanation"`
	ImagePath   string `json:"image_path"`
}

func AddQuestion(input QuestionInput) (*models.Question, error) {
	question := models.Question{
		PaperID:     input.PaperID,
		Content:     input.Content,
		Type:        input.Type,
		Options:     input.Options,
		Answer:      input.Answer,
		Explanation: input.Explanation,
		ImagePath:   input.ImagePath,
	}

	if err := database.DB.Create(&question).Error; err != nil {
		return nil, err
	}

	return &question, nil
}

func UpdateQuestionImage(questionID uint, imagePath string) (*models.Question, error) {
	var question models.Question
	if err := database.DB.First(&question, questionID).Error; err != nil {
		return nil, err
	}

	question.ImagePath = imagePath
	if err := database.DB.Save(&question).Error; err != nil {
		return nil, err
	}

	return &question, nil
}

func GetQuestionsByPaper(paperID uint) ([]models.Question, error) {
	var questions []models.Question
	if err := database.DB.Where("paper_id = ?", paperID).Order("id ASC").Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

func GetAllQuestions() ([]models.Question, error) {
	var questions []models.Question
	if err := database.DB.Order("created_at DESC").Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

func SearchQuestions(keyword string, paperID uint, qtype string, page, perPage int) ([]models.Question, int64, error) {
	var questions []models.Question
	var total int64

	query := database.DB.Model(&models.Question{})

	if keyword != "" {
		query = query.Where("content LIKE ? OR options LIKE ? OR answer LIKE ? OR explanation LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}
	if paperID > 0 {
		query = query.Where("paper_id = ?", paperID)
	}
	if qtype != "" {
		query = query.Where("type = ?", qtype)
	}

	query.Count(&total)

	offset := (page - 1) * perPage
	if err := query.Order("created_at DESC").Offset(offset).Limit(perPage).Find(&questions).Error; err != nil {
		return nil, 0, err
	}

	return questions, total, nil
} 