package user

import (
	"github.com/MuhammadIbraAlfathar/app-tweet/internal/schema"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *schema.User) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(user *schema.User) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		return nil
	})
}
