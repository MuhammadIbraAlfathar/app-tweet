package auth

import (
	"github.com/MuhammadIbraAlfathar/app-tweet/internal/domain/user"
	"github.com/MuhammadIbraAlfathar/app-tweet/internal/schema"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UseCase struct {
	userRepo user.Repository
}

func NewUseCase(userRepo user.Repository) *UseCase {
	return &UseCase{
		userRepo: userRepo,
	}
}

func (uc *UseCase) Register(req *RegisterRequest) error {
	//Hash&salt password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password: ", err)
	}
	id, err := uuid.NewV7()
	if err != nil {
		log.Println("Error generating UUID: ", err)
	}

	userEntity := schema.User{
		Id:              id,
		Name:            req.Name,
		Email:           req.Email,
		IsEmailVerified: false,
		Password:        string(passwordHash),
		Role:            schema.RoleUser(req.Role),
		Gender:          schema.Gender(req.Gender),
		ImageURL:        "",
	}

	err = uc.userRepo.Create(&userEntity)
	if err != nil {
		return err
	}

	return nil

}
