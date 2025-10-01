package services

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

var allowedExtensions = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".gif":  true,
	".bmp":  true,
}

func SaveImage(file *multipart.FileHeader, folder string) (string, error) {
	// 检查文件扩展名
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[ext] {
		return "", fmt.Errorf("不支持的文件格式")
	}

	// 生成唯一文件名
	uniqueID := uuid.New().String()[:8]
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("%s_%s%s", timestamp, uniqueID, ext)

	// 创建保存目录
	uploadFolder := filepath.Join("static", folder)
	if err := os.MkdirAll(uploadFolder, 0755); err != nil {
		return "", err
	}

	// 完整的文件保存路径
	filePath := filepath.Join(uploadFolder, filename)

	// 保存文件
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return "", err
	}

	// 返回相对路径
	relativePath := fmt.Sprintf("%s/%s", folder, filename)
	return relativePath, nil
}

func DeleteImage(imagePath string) bool {
	if imagePath == "" {
		return false
	}

	fullPath := filepath.Join("static", imagePath)
	if _, err := os.Stat(fullPath); err == nil {
		if err := os.Remove(fullPath); err == nil {
			return true
		}
	}
	return false
} 