package controller

import (
	"donasi-yuk/donasi"
	"donasi-yuk/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DonasiController struct {
	donasiService donasi.Service
}

func NewDonasiController(service donasi.Service) *DonasiController {
	return &DonasiController{donasiService: service}
}

// POST /api/donasi
func (h *DonasiController) CreateDonasi(c *gin.Context) {
	var input model.InputDonasi
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid!"})
		return
	}
	userID := c.MustGet("userID").(int)
	newDonasi, err := h.donasiService.CreateDonasi(input, uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat donasi!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": newDonasi})
}

// GET /api/donasi/user
func (h *DonasiController) GetDonasiByUser(c *gin.Context) {
	userID := uint(c.MustGet("userID").(int))
	//c.MustGet("userID").(int)
	donasii, err := h.donasiService.GetDonasiByUser(userID)
	//h.donasiService.GetDonasiByUser(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil donasi!"})
		return
	}
	formatted := donasi.FormatDonasis(donasii)
	c.JSON(http.StatusOK, gin.H{"data": formatted})
}

// GET /api/donasi/campaign/:id
func (h *DonasiController) GetDonasiByCampaign(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID campaign tidak valid"})
		return
	}
	donasii, err := h.donasiService.GetDonasiByCampaign(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil donasi campaign!"})
		return
	}
	formatted := donasi.FormatDonasis(donasii)
	c.JSON(http.StatusOK, gin.H{"data": formatted})
}

// GET /api/donasi
func (h *DonasiController) GetAllDonasi(c *gin.Context) {
	donasii, err := h.donasiService.GetAllDonasi()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil semua donasi!"})
		return
	}
	formatted := donasi.FormatDonasis(donasii)
	c.JSON(http.StatusOK, gin.H{"data": formatted})
}
