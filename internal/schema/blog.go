package schema

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Tags string

const (
	Tech        Tags = "Technology"
	Programming Tags = "Programming"
	Lifestyle   Tags = "Lifestyle"
	Business    Tags = "Business"
)

type Blog struct {
	Id        uuid.UUID      `json:"id" gorm:"primaryKey"`
	UserId    uuid.UUID      `json:"user_id" gorm:"not null"`
	Title     string         `json:"title" gorm:"type:varchar(250);not null"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	Category  string         `json:"category" gorm:"type:varchar(250);not null"`
	Tags      Tags           `json:"tags" gorm:"type:tags;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
