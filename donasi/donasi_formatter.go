package donasi

import (
	"donasi-yuk/model"
)

type DonasiResponse struct {
	ID            uint   `json:"id"`
	Amount        int    `json:"amount"`
	DonaturName   string `json:"donatur_name"`
	CampaignTitle string `json:"campaign_title"`
	CreatedAt     string `json:"created_at"`
}

func FormatDonasi(d model.Donasi) DonasiResponse {
	return DonasiResponse{
		ID:            d.ID,
		Amount:        d.Amount,
		DonaturName:   d.User.Name,
		CampaignTitle: d.Campaign.Title,
		CreatedAt:     d.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func FormatDonasis(donasis []model.Donasi) []DonasiResponse {
	var result []DonasiResponse
	for _, d := range donasis {
		result = append(result, FormatDonasi(d))
	}
	return result
}
