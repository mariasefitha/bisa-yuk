package controller

import (
	"donasi-yuk/campaign"
	"donasi-yuk/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignController struct {
	campaignService campaign.Service
}

func NewCampaignController(service campaign.Service) *CampaignController {
	return &CampaignController{campaignService: service}
}

// GET /campaigns
func (h *CampaignController) GetCampaigns(c *gin.Context) {
	// Ambil user ID dari JWT context (hasil middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Bisa dipakai untuk validasi, log, atau filter campaign milik user
	_ = userID // sementara kita belum pakai

	campaigns, err := h.campaignService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data campaign"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": campaigns})
}

// GET /campaigns/:id
func (h *CampaignController) GetCampaign(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	campaign, err := h.campaignService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": campaign})
}

// POST /campaigns
func (h *CampaignController) CreateCampaign(c *gin.Context) {
	var input model.CampaignInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := uint(c.MustGet("userID").(int)) // ✅ fix type conversion

	newCampaign, err := h.campaignService.Create(input, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat campaign"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": newCampaign})

	// Ambil user dari context (middleware auth)
	//userData := c.MustGet("userID").(model.User)

	// newCampaign, err := h.campaignService.Create(input, c.MustGet("userID").(uint))

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat campaign"})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"data": newCampaign})
}

// PUT /campaigns/:id
func (h *CampaignController) UpdateCampaign(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input model.CampaignInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userData := c.MustGet("userID").(model.User)

	updatedCampaign, err := h.campaignService.Update(id, input, userData.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updatedCampaign})
}

// DELETE /campaigns/:id
func (h *CampaignController) DeleteCampaign(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userData := c.MustGet("userID").(uint)

	err := h.campaignService.Delete(id, userData)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Campaign berhasil dihapus"})
}
