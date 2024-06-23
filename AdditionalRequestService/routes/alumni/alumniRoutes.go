package alumni

import (
	"AdditionalRequestService/controllers"
	"github.com/gin-gonic/gin"
	"shared/middleware"
)

func SetupRouter(route *gin.Engine) {
	protected := route.Group("/pass_requests")

	protected.Use(middleware.RequireAlumni)
	{
		protected.POST("", controllers.CreatePassRequest)
		protected.GET("", controllers.GetCurrentUserRequests)
		protected.PATCH("/:id", controllers.UpdatePassRequest)
		protected.DELETE("/:id", controllers.DeletePassRequest)
	}
}
