package campaign

import (
	"donasi-yuk/model"
	"errors"
)

type Service interface {
	FindAll() ([]model.Campaign, error)
	FindByID(id int) (model.Campaign, error)
	Create(input model.CampaignInput, userID uint) (model.Campaign, error)
	Update(id int, input model.CampaignInput, userID uint) (model.Campaign, error)
	Delete(id int, userID uint) error
}

// struct yang benar
type campaignService struct {
	repo Repository
}

// fungsi constructor
func NewService(repo Repository) Service {
	return &campaignService{repo}
}

// implementasi method-method interface
func (s *campaignService) FindAll() ([]model.Campaign, error) {
	return s.repo.FindAll()
}

func (s *campaignService) FindByID(id int) (model.Campaign, error) {
	campaign, err := s.repo.FindByID(id)
	if err != nil {
		return campaign, errors.New("campaign tidak ditemukan")
	}
	return campaign, nil
}

func (s *campaignService) Create(input model.CampaignInput, userID uint) (model.Campaign, error) {
	campaign := model.Campaign{
		Title:        input.Title,
		Description:  input.Description,
		TargetAmount: input.TargetAmount,
		ImageURL:     input.ImageURL,
		UserID:       userID,
	}
	return s.repo.Save(campaign)
}

func (s *campaignService) Update(id int, input model.CampaignInput, userID uint) (model.Campaign, error) {
	campaign, err := s.repo.FindByID(id)
	if err != nil {
		return campaign, errors.New("campaign tidak ditemukan")
	}

	if campaign.UserID != userID {
		return campaign, errors.New("kamu tidak berhak mengubah campaign ini")
	}

	campaign.Title = input.Title
	campaign.Description = input.Description
	campaign.TargetAmount = input.TargetAmount
	campaign.ImageURL = input.ImageURL

	return s.repo.Update(campaign)
}

func (s *campaignService) Delete(id int, userID uint) error {
	campaign, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("campaign tidak ditemukan")
	}

	if campaign.UserID != userID {
		return errors.New("tidak berhak menghapus campaign ini")
	}

	return s.repo.Delete(campaign)
}
