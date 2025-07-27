package model

import "time"

type Donasi struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `json:"user_id"`
	CampaignID uint      `json:"campaign_id"`
	Amount     int       `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	User     User     `gorm:"foreignKey:UserID" json:"user"`
	Campaign Campaign `gorm:"foreignKey:CampaignID" json:"campaign"`
}

type InputDonasi struct {
	CampaignID uint `json:"campaign_id" binding:"required"`
	Amount     int  `json:"amount" binding:"required"`
}
