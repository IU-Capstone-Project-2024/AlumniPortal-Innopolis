package main

import (
	"AuthService/routes/admin"
	"AuthService/routes/user"
	"alumniportal.com/shared/initializers"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
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
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type", "access-control-allow-origin", "access-control-allow-headers"},
	}))


	r.ForwardedByClientIP = true
	if r.SetTrustedProxies([]string{"127.0.0.1"}) != nil {
		panic("SetTrustedProxies failed")
	}

	user.SetupRouter(r)
	admin.SetupRouter(r)

	logrus.Info("Starting Auth Service")

	certFile := "certs/selfsigned.crt"
	keyFile := "certs/selfsigned.key"

	httpsServer := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	if err := httpsServer.ListenAndServeTLS(certFile, keyFile); err != nil && err != http.ErrServerClosed {
		logrus.Fatal("Failed to start HTTPS Auth Service:", err)
		return
	}
	logrus.Info("HTTPS Auth Service started")

	httpPort := ":8091"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("Redirecting request from %s to https://%s:8081%s", r.Host, r.Host, r.RequestURI)
		http.Redirect(w, r, "https://"+r.Host+":8081"+r.RequestURI, http.StatusMovedPermanently)
	})
	if err := http.ListenAndServe(httpPort, nil); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logrus.Fatal("Failed to start HTTP server for Auth Service:", err)
		panic(err)
	}
	logrus.Info("HTTP Auth Service started")
}
