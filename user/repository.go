package user

import (
	"donasi-yuk/model"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindByID(id int) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Save(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *repository) FindByID(id int) (model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return user, err
}
