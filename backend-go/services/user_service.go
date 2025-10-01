package services

import (
	"backend-go/database"
	"backend-go/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username, password, email, role string) (*models.User, error) {
	var existingUser models.User
	if err := database.DB.Where("username = ?", username).First(&existingUser).Error; err == nil {
		return nil, errors.New("用户名已存在")
	}

	if email != "" {
		if err := database.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
			return nil, errors.New("邮箱已存在")
		}
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Username:     username,
		PasswordHash: string(passwordHash),
		Email:        email,
		Role:         role,
		IsActive:     true,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func VerifyUser(username, password string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	return &user, nil
}

func GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
} 