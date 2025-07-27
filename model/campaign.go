package model

import "gorm.io/gorm"

type Campaign struct {
	gorm.Model
	Title        string `json:"title"`
	Description  string `json:"description"`
	TargetAmount int    `json:"target_amount"`
	ImageURL     string `json:"image_url"`
	UserID       uint   `json:"user_id"` // foreign key ke User
}

//INPUT
type CampaignInput struct {
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description" binding:"required"`
	TargetAmount int    `json:"target_amount" binding:"required"`
	ImageURL     string `json:"image_url"`
}
