package handlers

import (
	"backend-go/database"
	"backend-go/models"
	"backend-go/services"
	"backend-go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddPaper(c *gin.Context) {
	var input services.PaperInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, 400, utils.PARAMERR, "参数错误")
		return
	}

	paper, err := services.AddPaper(input)
	if err != nil {
		utils.Error(c, 500, utils.DBERR, "添加试卷失败")
		return
	}

	utils.SuccessWithMsg(c, "试卷添加成功", gin.H{"id": paper.ID})
}

func UpdatePaperImage(c *gin.Context) {
	paperID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Error(c, 400, utils.PARAMERR, "无效的试卷ID")
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

	_, err = services.UpdatePaperImage(uint(paperID), req.ImagePath)
	if err != nil {
		utils.Error(c, 404, utils.NODATA, "试卷不存在")
		return
	}

	utils.SuccessWithMsg(c, "试卷图片更新成功", nil)
}

func GetPapers(c *gin.Context) {
	course := c.Query("course")
	yearStr := c.Query("year")
	term := c.Query("term")
	college := c.Query("college")
	keyword := c.Query("keyword")
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("per_page", "10")

	page, _ := strconv.Atoi(pageStr)
	perPage, _ := strconv.Atoi(perPageStr)
	year, _ := strconv.Atoi(yearStr)

	papers, total, err := services.GetPapers(course, term, college, keyword, year, page, perPage)
	if err != nil {
		utils.Error(c, 500, utils.DBERR, "查询失败")
		return
	}

	if len(papers) == 0 {
		utils.Error(c, 404, utils.NODATA, "暂无试卷数据")
		return
	}

	result := make([]gin.H, 0, len(papers))
	for _, p := range papers {
		result = append(result, gin.H{
			"id":         p.ID,
			"course":     p.Course,
			"year":       p.Year,
			"term":       p.Term,
			"college":    p.College,
			"image_path": p.ImagePath,
			"created_at": p.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	utils.Success(c, gin.H{
		"total": total,
		"items": result,
	})
}

func GetStats(c *gin.Context) {
	var paperCount, questionCount, userCount int64

	database.DB.Model(&models.Paper{}).Count(&paperCount)
	database.DB.Model(&models.Question{}).Count(&questionCount)
	database.DB.Model(&models.User{}).Count(&userCount)

	utils.Success(c, gin.H{
		"papers":    paperCount,
		"questions": questionCount,
		"users":     userCount,
	})
} 