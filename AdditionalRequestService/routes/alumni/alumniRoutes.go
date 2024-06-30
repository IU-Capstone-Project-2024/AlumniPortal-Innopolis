package alumni

import (
	"AdditionalRequestService/controllers"
	"alumniportal.com/shared/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	protected := route.Group("/requests")

	protected.Use(middleware.RequireAlumni)
	{
		protected.POST("", controllers.CreatePassRequest)
		protected.GET("", controllers.GetCurrentUserRequests)
		protected.GET("/:id", controllers.GetPassRequest)
		protected.PATCH("/:id", controllers.UpdatePassRequest)
		protected.DELETE("/:id", controllers.DeletePassRequest)
	}
}
