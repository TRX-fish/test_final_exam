package handlers

import (
	"backend-go/middleware"
	"backend-go/services"
	"backend-go/utils"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, utils.PARAMERR, "参数错误")
		return
	}

	if req.Username == "" || req.Password == "" {
		utils.Error(c, 400, utils.PARAMERR, "用户名和密码不能为空")
		return
	}

	user, err := services.RegisterUser(req.Username, req.Password, req.Email, "user")
	if err != nil {
		utils.Error(c, 409, utils.DATAEXIST, err.Error())
		return
	}

	utils.SuccessWithMsg(c, "注册成功", gin.H{
		"id":       user.ID,
		"username": user.Username,
	})
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, utils.PARAMERR, "参数错误")
		return
	}

	if req.Username == "" || req.Password == "" {
		utils.Error(c, 400, utils.PARAMERR, "用户名和密码不能为空")
		return
	}

	user, err := services.VerifyUser(req.Username, req.Password)
	if err != nil {
		utils.Error(c, 401, utils.LOGINERR, "用户名或密码错误")
		return
	}

	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		utils.Error(c, 500, utils.INTERNALERR, "生成token失败")
		return
	}

	utils.SuccessWithMsg(c, "登录成功", gin.H{
		"token":   token,
		"user_id": user.ID,
		"role":    user.Role,
	})
}

func GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Error(c, 401, utils.SESSIONERR, "用户未登录")
		return
	}

	user, err := services.GetUserByID(userID.(uint))
	if err != nil {
		utils.Error(c, 404, utils.USERERR, "用户不存在")
		return
	}

	utils.Success(c, gin.H{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"role":       user.Role,
		"is_active":  user.IsActive,
		"created_at": user.CreatedAt.Format("2006-01-02 15:04:05"),
	})
} 