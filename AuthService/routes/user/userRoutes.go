package user

import (
	"AuthService/controllers"
	"alumniportal.com/shared/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	public := route.Group("/auth")
	{
		public.POST("signup", controllers.Signup)
		public.POST("login", controllers.Login)
	}

	protected := route.Group("/auth")
	protected.Use(middleware.RequireAuth)
	{
		protected.GET("validate", controllers.Validate)
		protected.GET("user", controllers.GetInfo)
		protected.POST("logout", controllers.Logout)
	}
}
