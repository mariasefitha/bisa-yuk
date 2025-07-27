package donasi

import (
	"donasi-yuk/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(donasi model.Donasi) (model.Donasi, error)
	FindByUserID(userID uint) ([]model.Donasi, error)
	FindByCampaignID(campaignID uint) ([]model.Donasi, error)
	FindAll() ([]model.Donasi, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(donasi model.Donasi) (model.Donasi, error) {
	err := r.db.Create(&donasi).Error
	return donasi, err
}

func (r *repository) FindByUserID(userID uint) ([]model.Donasi, error) {
	var donasii []model.Donasi
	err := r.db.Preload("Campaign").Where("user_id = ?", userID).Find(&donasii).Error
	return donasii, err
}

// Find Campaign ID
func (r *repository) FindByCampaignID(campaignID uint) ([]model.Donasi, error) {
	var donasii []model.Donasi
	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Find(&donasii).Error
	return donasii, err
}

// Find All Donasi
func (r *repository) FindAll() ([]model.Donasi, error) {
	var donasii []model.Donasi
	err := r.db.Preload("User").Preload("Campaign").Find(&donasii).Error
	return donasii, err
}
