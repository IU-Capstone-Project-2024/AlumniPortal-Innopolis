package main

import (
	"AdditionalRequestService/routes/admin"
	"AdditionalRequestService/routes/alumni"
	"alumniportal.com/shared/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://alumni-inno.netlify.app"},
		AllowMethods:     []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "access-control-allow-origin", "access-control-allow-headers"},
		AllowCredentials: true,
	}))

	// r.Use(middleware.AuthenticateToken())

	r.ForwardedByClientIP = true
	if r.SetTrustedProxies([]string{"127.0.0.1"}) != nil {
		panic("SetTrustedProxies failed")
	}

	alumni.SetupRouter(r)
	admin.SetupRouter(r)

	logrus.Info("Starting Request Service")

	if err := r.Run(":8082"); err != nil {
		logrus.Fatal("Failed to start Request Service")
		panic(err)
	}
}
