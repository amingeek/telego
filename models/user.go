package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;size:100;not null" json:"email"`
	Phone    string `gorm:"uniqueIndex;size:15" json:"phone,omitempty"`

	DisplayName string `gorm:"size:100" json:"display_name"`
	Bio         string `gorm:"size:500" json:"bio,omitempty"`
	AvatarURL   string `gorm:"size:255" json:"avatar_url,omitempty"`

	PasswordHash string `gorm:"not null" json:"-"`

	IsActive   bool       `gorm:"default:true" json:"is_active"`
	IsPrivate  bool       `gorm:"default:false" json:"is_private"`
	LastSeenAt *time.Time `json:"last_seen_at,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
