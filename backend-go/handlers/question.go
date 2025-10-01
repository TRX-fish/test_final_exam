package handlers

import (
	"backend-go/services"
	"backend-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddQuestion(c *gin.Context) {
	var input services.QuestionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, 400, utils.PARAMERR, "参数错误")
		return
	}

	question, err := services.AddQuestion(input)
	if err != nil {
		utils.Error(c, 500, utils.DBERR, "添加题目失败")
		return
	}

	utils.SuccessWithMsg(c, "题目添加成功", gin.H{"id": question.ID})
}

func UpdateQuestionImage(c *gin.Context) {
	questionID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, 400, utils.PARAMERR, "无效的题目ID")
		return
	}

	var req struct {
		ImagePath string `json:"image_path"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, utils.PARAMERR, "参数错误")
		return
	}

	if req.ImagePath == "" {
		utils.Error(c, 400, utils.PARAMERR, "图片路径不能为空")
		return
	}

	_, err = services.UpdateQuestionImage(uint(questionID), req.ImagePath)
	if err != nil {
		utils.Error(c, 404, utils.NODATA, "题目不存在")
		return
	}

	utils.SuccessWithMsg(c, "题目图片更新成功", nil)
}

func GetQuestions(c *gin.Context) {
	paperIDStr := c.Query("paper_id")

	if paperIDStr != "" {
		paperID, _ := strconv.ParseUint(paperIDStr, 10, 32)
		questions, err := services.GetQuestionsByPaper(uint(paperID))
		if err != nil || len(questions) == 0 {
			utils.Error(c, 404, utils.NODATA, "暂无题目数据")
			return
		}

		result := make([]gin.H, 0, len(questions))
		for _, q := range questions {
			result = append(result, gin.H{
				"id":          q.ID,
				"paper_id":    q.PaperID,
				"content":     q.Content,
				"type":        q.Type,
				"options":     q.Options,
				"answer":      q.Answer,
				"explanation": q.Explanation,
				"image_path":  q.ImagePath,
				"created_at":  q.CreatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		utils.Success(c, result)
	} else {
		questions, err := services.GetAllQuestions()
		if err != nil || len(questions) == 0 {
			utils.Error(c, 404, utils.NODATA, "暂无题目数据")
			return
		}

		result := make([]gin.H, 0, len(questions))
		for _, q := range questions {
			result = append(result, gin.H{
				"id":          q.ID,
				"paper_id":    q.PaperID,
				"content":     q.Content,
				"type":        q.Type,
				"options":     q.Options,
				"answer":      q.Answer,
				"explanation": q.Explanation,
				"image_path":  q.ImagePath,
				"created_at":  q.CreatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		utils.Success(c, result)
	}
}

func SearchQuestions(c *gin.Context) {
	keyword := c.Query("keyword")
	paperIDStr := c.Query("paper_id")
	qtype := c.Query("type")
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("per_page", "10")

	page, _ := strconv.Atoi(pageStr)
	perPage, _ := strconv.Atoi(perPageStr)
	paperID, _ := strconv.ParseUint(paperIDStr, 10, 32)

	questions, total, err := services.SearchQuestions(keyword, uint(paperID), qtype, page, perPage)
	if err != nil {
		utils.Error(c, 500, utils.DBERR, "查询失败")
		return
	}

	if len(questions) == 0 {
		utils.Error(c, 404, utils.NODATA, "暂无题目数据")
		return
	}

	result := make([]gin.H, 0, len(questions))
	for _, q := range questions {
		result = append(result, gin.H{
			"id":          q.ID,
			"paper_id":    q.PaperID,
			"content":     q.Content,
			"type":        q.Type,
			"options":     q.Options,
			"answer":      q.Answer,
			"explanation": q.Explanation,
			"image_path":  q.ImagePath,
			"created_at":  q.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	utils.Success(c, gin.H{
		"total": total,
		"items": result,
	})
} 