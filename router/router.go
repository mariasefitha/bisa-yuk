package router

import (
	"donasi-yuk/controller"
	"donasi-yuk/middleware"
	"donasi-yuk/user"

	"github.com/gin-gonic/gin"
)

// func SetupRoutes(r *gin.Engine, userController *controller.UserController) {
// 	// Public routes
// 	r.POST("/register", userController.Register)
// 	r.POST("/login", userController.Login)

// 	// Protected routes
// 	auth := r.Group("/")
// 	auth.Use(middleware.AuthMiddleware())
// 	// auth.GET("/campaigns", campaignController.GetCampaigns)
// }

func SetupRoutes(r *gin.Engine, userController *controller.UserController, campaignController *controller.CampaignController, donasiController *controller.DonasiController, userService user.Service) {
	// USER
	user := r.Group("/api/users")
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
	}

	// CAMPAIGN
	campaign := r.Group("/api/campaigns")
	campaign.Use(middleware.AuthMiddleware())
	{
		campaign.GET("/", campaignController.GetCampaigns)
		campaign.GET("/:id", campaignController.GetCampaign)
		campaign.POST("/", campaignController.CreateCampaign)
		campaign.PUT("/:id", campaignController.UpdateCampaign)
		campaign.DELETE("/:id", campaignController.DeleteCampaign)
	}

	//DONASI
	donasiRoutes := r.Group("/api/donasi")
	donasiRoutes.Use(middleware.AuthMiddleware())
	{
		donasiRoutes.POST("/", donasiController.CreateDonasi)
		donasiRoutes.GET("/user", donasiController.GetDonasiByUser)
		donasiRoutes.GET("/campaign/:id", donasiController.GetDonasiByCampaign)
		donasiRoutes.GET("/", donasiController.GetAllDonasi)
	}
}
