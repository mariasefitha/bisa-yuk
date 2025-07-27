package service

import (
	"donasi-yuk/model"
	"donasi-yuk/user"
	"errors"
)

type UserService interface {
	Register(input model.RegisterInput) (model.User, error)
	Login(input model.LoginInput) (model.User, error)
	GetUserByID(id int) (model.User, error)
}

// Fungsi Get User By ID
//
//	func (u *userService) GetUserByID(id int) (model.User, error) {
//		return model.User{}, errors.New("fitur get user by ID belum diimplementasi")
//	}
func (s *userService) GetUserByID(id int) (model.User, error) {
	return s.repo.FindByID(id)
}

// Fungsi Login
func (s *userService) Login(input model.LoginInput) (model.User, error) {
	// Cek apakah user dengan email tersebut ada
	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		return model.User{}, err
	}

	if user.ID == 0 {
		return model.User{}, errors.New("User tidak ditemukan atau belum ter-registrasi")
	}

	// Cek apakah password cocok (tanpa hash, plain-text)
	if user.Password != input.Password {
		return model.User{}, errors.New("Passwordnya salah ges!")
	}

	return user, nil
}

// Register implements user.Service.
func (s *userService) Register(input model.RegisterInput) (model.User, error) {
	// Cek apakah email sudah terdaftar
	existingUser, _ := s.repo.FindByEmail(input.Email)
	if existingUser.ID != 0 {
		return model.User{}, errors.New("Email sudah terdafta di aplikasi!")
	}

	// Simpan user baru
	newUser := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password, // optional: kamu bisa hash password di sini
	}

	createdUser, err := s.repo.Save(newUser)
	if err != nil {
		return model.User{}, err
	}

	return createdUser, nil
}

type userService struct {
	repo user.Repository
}

func NewService(repo user.Repository) UserService {
	return &userService{repo}
}

// func (s *userService) Register(input model.RegisterInput) (model.User, error) {
// 	existingUser, _ := s.repo.FindByEmail(input.Email)
// 	if existingUser.ID != 0 {
// 		return model.User{}, errors.New("Email sudah terdaftar!")
// 	}

// 	newUser := model.User{
// 		Name:     input.Name,
// 		Email:    input.Email,
// 		Password: input.Password,
// 	}

// 	createdUser, err := s.repo.Save(newUser)
// 	if err != nil {
// 		return model.User{}, err
// 	}

// 	return createdUser, nil
// }

// func (s *userService) Login(input model.LoginInput) (model.User, error) {
// 	foundUser, err := s.repo.FindByEmail(input.Email)
// 	if err != nil || foundUser.ID == 0 {
// 		return model.User{}, errors.New("User tidak ditemukan!")
// 	}

// 	if foundUser.Password != input.Password {
// 		return model.User{}, errors.New("Password salah!")
// 	}

// 	return foundUser, nil
// }
