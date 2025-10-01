package models

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;comment:用户ID，主键自增" json:"id"`
	Username     string    `gorm:"type:varchar(50);uniqueIndex;not null;comment:用户名，唯一" json:"username"`
	PasswordHash string    `gorm:"type:varchar(128);not null;comment:密码哈希值" json:"-"`
	Email        string    `gorm:"type:varchar(100);uniqueIndex;comment:邮箱，唯一，可选" json:"email"`
	Role         string    `gorm:"type:enum('user','admin');not null;default:user;comment:用户角色" json:"role"`
	IsActive     bool      `gorm:"not null;default:true;comment:账号是否激活" json:"is_active"`
	CreatedAt    time.Time `gorm:"comment:注册时间" json:"created_at"`
	UpdatedAt    time.Time `gorm:"comment:信息更新时间" json:"updated_at"`
}

func (User) TableName() string {
	return "user"
} 