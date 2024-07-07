package main

import (
	"DonationService/routes/admin"
	"DonationService/routes/alumni"
	"DonationService/routes/student"

	"alumniportal.com/shared/initializers"
	"alumniportal.com/shared/middleware"
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
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	r.Use(middleware.AuthenticateToken())

	r.ForwardedByClientIP = true
	if r.SetTrustedProxies([]string{"127.0.0.1"}) != nil {
		panic("SetTrustedProxies failed")
	}

	student.SetupRouter(r)
	admin.SetupRouter(r)
	alumni.SetupRouter(r)

	err := r.Run(":3000")

	if err != nil {
		panic("Error starting Event Service")
	}
	logrus.Info("Event Service successfully started!")
}
