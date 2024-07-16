package student

import (
	"VolunteerService/controllers"

	"alumniportal.com/shared/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	protected := route.Group("/volunteer")

	protected.Use(middleware.RequireStudent)
	{
		protected.GET("/unverified", controllers.GetUnverifiedVolunteers)
		protected.GET("/unverified/:id", controllers.GetStudentVolunteerRequest)
		protected.POST("/unverified/:id/approve", controllers.AcceptVolunteerRequest)
		protected.POST("/unverified/:id/decline", controllers.DeclineVolunteerRequest)
	}
}
