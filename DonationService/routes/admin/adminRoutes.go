package admin

import (
	"DonationService/controllers"

	"alumniportal.com/shared/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	admin := route.Group("/donation")
	admin.Use(middleware.RequireAdminRights)
	{
		admin.GET("/user_id", controllers.GetCurrentUserDonationRequests)
		admin.GET("/amount", controllers.GetCurrentAmountDonationRequests)
		admin.GET("/date", controllers.GetCurrentDateDonationRequests)
	}
}
