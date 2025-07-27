package main

import (
	"donasi-yuk/campaign"
	"donasi-yuk/config"
	"donasi-yuk/controller"
	"donasi-yuk/donasi"
	"donasi-yuk/router"
	"donasi-yuk/service"
	"donasi-yuk/user"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	db := config.DB

	// USER setup
	userRepo := user.NewRepository(db)
	userService := service.NewService(userRepo)
	userController := controller.NewUserController(userService)

	// CAMPAIGN setup
	campaignRepo := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepo)
	campaignController := controller.NewCampaignController(campaignService)

	//DONASI setup
	donasiRepo := donasi.NewRepository(db)
	donasiService := donasi.NewService(donasiRepo)
	donasiController := controller.NewDonasiController(donasiService)

	// ROUTING
	r := gin.Default()
	router.SetupRoutes(r, userController, campaignController, donasiController, userService)

	// Run
	r.Run(":8080")
}
