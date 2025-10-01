package handlers

import (
	"backend-go/services"
	"backend-go/utils"

	"github.com/gin-gonic/gin"
)

func UploadPaperImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.Error(c, 400, utils.PARAMERR, "没有选择文件")
		return
	}

	if file.Filename == "" {
		utils.Error(c, 400, utils.PARAMERR, "没有选择文件")
		return
	}

	imagePath, err := services.SaveImage(file, "paper_images")
	if err != nil {
		utils.Error(c, 500, utils.IOERR, err.Error())
		return
	}

	utils.SuccessWithMsg(c, "试卷图片上传成功", gin.H{"image_path": imagePath})
}

func UploadQuestionImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.Error(c, 400, utils.PARAMERR, "没有选择文件")
		return
	}

	if file.Filename == "" {
		utils.Error(c, 400, utils.PARAMERR, "没有选择文件")
		return
	}

	imagePath, err := services.SaveImage(file, "question_images")
	if err != nil {
		utils.Error(c, 500, utils.IOERR, err.Error())
		return
	}

	utils.SuccessWithMsg(c, "题目图片上传成功", gin.H{"image_path": imagePath})
}

func DeleteImage(c *gin.Context) {
	var req struct {
		ImagePath  string `json:"image_path"`
		PaperID    uint   `json:"paper_id"`
		QuestionID uint   `json:"question_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, utils.PARAMERR, "参数错误")
		return
	}

	if req.ImagePath == "" {
		utils.Error(c, 400, utils.PARAMERR, "图片路径不能为空")
		return
	}

	fileDeleted := services.DeleteImage(req.ImagePath)
	dbUpdated := false
	msg := ""

	if req.PaperID > 0 {
		_, err := services.UpdatePaperImage(req.PaperID, "")
		dbUpdated = err == nil
		msg = "试卷图片"
	} else if req.QuestionID > 0 {
		_, err := services.UpdateQuestionImage(req.QuestionID, "")
		dbUpdated = err == nil
		msg = "题目图片"
	} else {
		msg = "图片"
	}

	if fileDeleted {
		if dbUpdated || (req.PaperID == 0 && req.QuestionID == 0) {
			utils.SuccessWithMsg(c, msg+"删除成功", nil)
		} else {
			utils.Error(c, 500, utils.DBERR, msg+"数据库字段清空失败")
		}
	} else {
		utils.Error(c, 500, utils.IOERR, msg+"删除失败")
	}
} 