package services

import (
	"backend-go/config"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// Storage 接口定义
type Storage interface {
	SaveFile(file *multipart.FileHeader, folder string) (string, error)
	DeleteFile(filePath string) error
	GetFileURL(filePath string) string
}

// 当前使用的存储类型（可通过配置文件控制）
var currentStorage Storage

func InitStorage() {
	// 根据配置选择存储方式
	if config.AppConfig.Storage.Type == "oss" {
		currentStorage = NewOSSStorage(
			config.AppConfig.OSS.Endpoint,
			config.AppConfig.OSS.AccessKeyID,
			config.AppConfig.OSS.AccessKeySecret,
			config.AppConfig.OSS.BucketName,
		)
	} else {
		currentStorage = NewLocalStorage()
	}
}

// GetStorage 获取当前存储服务
func GetStorage() Storage {
	return currentStorage
}

// ============= 本地存储实现 =============

type LocalStorage struct {
	baseURL string
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		baseURL: "http://47.93.252.105:5000", // 可从配置读取
	}
}

func (s *LocalStorage) SaveFile(file *multipart.FileHeader, folder string) (string, error) {
	return SaveImage(file, folder) // 复用现有函数
}

func (s *LocalStorage) DeleteFile(filePath string) error {
	fullPath := filepath.Join("static", filePath)
	return os.Remove(fullPath)
}

func (s *LocalStorage) GetFileURL(filePath string) string {
	// 返回完整URL
	return fmt.Sprintf("%s/static/%s", s.baseURL, filePath)
}

// ============= OSS存储实现（未来使用）=============

type OSSStorage struct {
	endpoint        string
	accessKeyID     string
	accessKeySecret string
	bucketName      string
	baseURL         string
}

func NewOSSStorage(endpoint, keyID, keySecret, bucket string) *OSSStorage {
	return &OSSStorage{
		endpoint:        endpoint,
		accessKeyID:     keyID,
		accessKeySecret: keySecret,
		bucketName:      bucket,
		baseURL:         config.AppConfig.OSS.BaseURL,
	}
}

func (s *OSSStorage) SaveFile(file *multipart.FileHeader, folder string) (string, error) {
	// 导入OSS SDK
	// import "github.com/aliyun/aliyun-oss-go-sdk/oss"
	// import "time"
	
	client, err := NewOSSClient()
	if err != nil {
		return "", err
	}
	
	bucket, err := client.Bucket(s.bucketName)
	if err != nil {
		return "", err
	}
	
	// 生成唯一文件名
	ext := filepath.Ext(file.Filename)
	uniqueID := fmt.Sprintf("%d", time.Now().UnixNano())
	objectKey := fmt.Sprintf("%s/%s%s", folder, uniqueID, ext)
	
	// 打开上传文件
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	
	// 上传到OSS
	err = bucket.PutObject(objectKey, src)
	if err != nil {
		return "", err
	}
	
	// 返回相对路径（前端会拼接baseURL）
	return objectKey, nil
}

func (s *OSSStorage) DeleteFile(filePath string) error {
	client, err := NewOSSClient()
	if err != nil {
		return err
	}
	
	bucket, err := client.Bucket(s.bucketName)
	if err != nil {
		return err
	}
	
	return bucket.DeleteObject(filePath)
}

// NewOSSClient 创建OSS客户端
func NewOSSClient() (*oss.Client, error) {
	return oss.New(
		config.AppConfig.OSS.Endpoint,
		config.AppConfig.OSS.AccessKeyID,
		config.AppConfig.OSS.AccessKeySecret,
	)
}

func (s *OSSStorage) GetFileURL(filePath string) string {
	// 返回OSS带签名的临时URL（24小时有效）
	client, err := NewOSSClient()
	if err != nil {
		return ""
	}

	bucket, err := client.Bucket(s.bucketName)
	if err != nil {
		return ""
	}

	// 生成24小时有效的签名URL
	signedURL, err := bucket.SignURL(filePath, oss.HTTPGet, 86400)
	if err != nil {
		return ""
	}

	return signedURL
}

// ============= 统一调用接口 =============

// UploadPaperFile 上传试卷文件（图片或PDF）
func UploadPaperFile(file *multipart.FileHeader) (string, error) {
	storage := GetStorage()
	return storage.SaveFile(file, "papers_pdf")
}

// DeletePaperFile 删除试卷文件
func DeletePaperFile(filePath string) error {
	storage := GetStorage()
	return storage.DeleteFile(filePath)
}

// GetPaperFileURL 获取试卷文件完整URL
func GetPaperFileURL(filePath string) string {
	if filePath == "" {
		return ""
	}
	storage := GetStorage()
	return storage.GetFileURL(filePath)
} 