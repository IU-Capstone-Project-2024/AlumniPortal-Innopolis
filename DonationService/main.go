package main

import (
	"DonationService/routes/admin"
	"DonationService/routes/alumni"
	"DonationService/routes/student"

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

	student.SetupRouter(r)
	admin.SetupRouter(r)
	alumni.SetupRouter(r)

	logrus.Info("Starting Donation Service")

	if err := r.Run(":8083"); err != nil {
		logrus.Fatal("Error starting Donation Service")
		panic(err)
	}
}
