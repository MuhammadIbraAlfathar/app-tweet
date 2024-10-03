package schema

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Gender string

const (
	FemaleGender Gender = "female"
	MaleGender   Gender = "male"
)

type RoleUser string

const (
	Admin RoleUser = "admin"
	Users RoleUser = "user"
	Guest RoleUser = "guest"
)

type User struct {
	Id              uuid.UUID      `json:"id" gorm:"primaryKey"`
	Email           string         `json:"email" gorm:"type:varchar(230);unique;not null;index:,type:hash"`
	IsEmailVerified bool           `json:"is_email_verified" gorm:"not null;default:false"`
	Name            string         `json:"name" gorm:"type:varchar(230);not null"`
	Password        string         `gorm:"type:varchar(230);not null"`
	Gender          Gender         `json:"gender" gorm:"type:gender_user;not null"`
	ImageURL        string         `json:"image_url" gorm:"type:text"`
	Role            RoleUser       `gorm:"type:users_role;not null;default:user"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
