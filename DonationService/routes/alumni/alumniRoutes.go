package alumni

import (
	"DonationService/controllers"

	"alumniportal.com/shared/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	alumni := route.Group("/donation")
	alumni.Use(middleware.RequireAlumni)
	{
		alumni.POST("", controllers.CreateDonationRequest)
		alumni.GET("", controllers.GetCurrentUserDonationRequests)
		alumni.GET("/:id", controllers.GetDonationRequest)
	}
}
