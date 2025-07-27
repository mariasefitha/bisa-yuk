package user

import (
	"donasi-yuk/config"
	"donasi-yuk/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input model.RegisterInput) (model.User, error)
	Login(input model.LoginInput) (model.User, error)
	GetUserByID(id int) (model.User, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

// Register new user
func (s *service) Register(input model.RegisterInput) (model.User, error) {
	var user model.User

	// Cek apakah email sudah digunakan
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err == nil {
		return user, errors.New("email sudah digunakan")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	// Buat user baru
	user = model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     "user", // default
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// Login
func (s *service) Login(input model.LoginInput) (model.User, error) {
	var user model.User

	// Cari user berdasarkan email
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return user, errors.New("email tidak ditemukan")
	}

	// Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, errors.New("password salah")
	}

	return user, nil
}

// Get user by ID (dipakai di middleware)
func (s *service) GetUserByID(id int) (model.User, error) {
	var user model.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}
