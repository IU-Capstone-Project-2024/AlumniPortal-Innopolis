package student

import (
	"DonationService/controllers"

	"alumniportal.com/shared/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	student := route.Group("/donation")
	student.Use(middleware.RequireStudent)
	{
		student.GET("/project/sum", controllers.GetAccumulatedSumDonationRequest)
	}
}
