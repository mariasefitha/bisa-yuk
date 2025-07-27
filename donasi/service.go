package donasi

import (
	"donasi-yuk/model"
)

type Service interface {
	CreateDonasi(input model.InputDonasi, userID uint) (model.Donasi, error)
	GetDonasiByUser(userID uint) ([]model.Donasi, error)
	GetDonasiByCampaign(campaignID uint) ([]model.Donasi, error)
	GetAllDonasi() ([]model.Donasi, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateDonasi(input model.InputDonasi, userID uint) (model.Donasi, error) {
	donasi := model.Donasi{
		UserID:     userID,
		CampaignID: input.CampaignID,
		Amount:     input.Amount,
	}

	createdDonasi, err := s.repo.Create(donasi)
	if err != nil {
		return model.Donasi{}, err
	}

	return createdDonasi, nil
}

// GET Donasi by User
func (s *service) GetDonasiByUser(userID uint) ([]model.Donasi, error) {
	return s.repo.FindByUserID(userID)
	// func (s *donasiService) GetDonasiByUser(userID uint) ([]model.Donasi, error) {
	// return s.repo.FindByUserID(userID) }
}

// GET Donasi by Campaign
func (s *service) GetDonasiByCampaign(campaignID uint) ([]model.Donasi, error) {
	return s.repo.FindByCampaignID(campaignID)
}

// GET All Donasi
func (s *service) GetAllDonasi() ([]model.Donasi, error) {
	return s.repo.FindAll()
}
