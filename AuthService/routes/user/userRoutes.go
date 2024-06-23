package user

import (
	"AuthService/controllers"
	"github.com/gin-gonic/gin"
	"shared/middleware"
)

func SetupRouter(route *gin.Engine) {
	public := route.Group("/")
	{
		public.POST("signup", controllers.Signup)
		public.POST("login", controllers.Login)
	}

	protected := route.Group("/")
	protected.Use(middleware.RequireAuth)
	{
		protected.GET("validate", controllers.Validate)
		protected.GET("user", controllers.GetInfo)
	}
}
