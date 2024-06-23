package AdditionalRequestService

import (
	"AdditionalRequestService/routes/admin"
	"AdditionalRequestService/routes/alumni"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"shared/initializers"
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
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	r.ForwardedByClientIP = true
	if r.SetTrustedProxies([]string{"127.0.0.1"}) != nil {
		panic("SetTrustedProxies failed")
	}

	alumni.SetupRouter(r)
	admin.SetupRouter(r)

	err := r.Run()

	if err != nil {
		panic("Error starting AuthService")
	}

}
