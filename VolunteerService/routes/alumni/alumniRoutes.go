package alumni

import (
	"VolunteerService/controllers"

	"alumniportal.com/shared/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	protected := route.Group("/volunteer")

	protected.Use(middleware.RequireAlumni)
	{
		protected.POST("", controllers.CreateVolunteerRequest)
		protected.GET("", controllers.GetCurrentUserRequests)
		protected.GET("/:id", controllers.GetVolunteerRequest)
		protected.DELETE("/:id/delete", controllers.DeleteVolunteerRequest)
	}
}
