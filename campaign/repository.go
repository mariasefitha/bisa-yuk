package campaign

import (
	"donasi-yuk/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]model.Campaign, error)
	FindByID(id int) (model.Campaign, error)
	Save(campaign model.Campaign) (model.Campaign, error)
	Update(campaign model.Campaign) (model.Campaign, error)
	Delete(campaign model.Campaign) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]model.Campaign, error) {
	var campaigns []model.Campaign
	err := r.db.Find(&campaigns).Error
	return campaigns, err
}

func (r *repository) FindByID(id int) (model.Campaign, error) {
	var campaign model.Campaign
	err := r.db.First(&campaign, id).Error
	return campaign, err
}

func (r *repository) Save(campaign model.Campaign) (model.Campaign, error) {
	err := r.db.Create(&campaign).Error
	return campaign, err
}

func (r *repository) Update(campaign model.Campaign) (model.Campaign, error) {
	err := r.db.Save(&campaign).Error
	return campaign, err
}

func (r *repository) Delete(campaign model.Campaign) error {
	err := r.db.Delete(&campaign).Error
	return err
}
