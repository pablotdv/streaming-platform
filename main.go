package main

import (
	"context"
	"log"
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pablotdv/streaming-platform/data"
	"github.com/pablotdv/streaming-platform/models"
	"github.com/pablotdv/streaming-platform/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	var err error
	data.Db, err = gorm.Open(mysql.Open("restreamer:restreamer@tcp(127.0.0.1:3306)/restreamer?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	data.Db.AutoMigrate(&models.Streamer{})
}

var provider *oidc.Provider
var verifier *oidc.IDTokenVerifier

func init() {
	ctx := context.Background()

	// Inicialize o provider OIDC usando o endpoint de descoberta do IdentityServer.
	var err error
	provider, err = oidc.NewProvider(ctx, "https://localhost:5001")
	if err != nil {
		log.Fatalf("Failed to get provider: %v", err)
	}

	// Configure um OIDC Config para o token.
	oidcConfig := &oidc.Config{
		ClientID: "streaming-api", // Substitua com o ID do cliente registrado no IdentityServer.
	}

	// Crie um verificador para tokens.
	verifier = provider.Verifier(oidcConfig)
}

func OIDCTokenValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawToken := c.Request.Header.Get("Authorization")
		if rawToken == "" {
			c.JSON(http.StatusUnauthorized, "Authorization header missing")
			c.Abort()
			return
		}

		// Remova o prefixo "Bearer" se estiver presente.
		if len(rawToken) > 7 && rawToken[0:7] == "Bearer " {
			rawToken = rawToken[7:]
		}

		_, err := verifier.Verify(c, rawToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
	}))
	api := router.Group("/api")
	{
		api.Use(OIDCTokenValidation())
		streamers := api.Group("/streamers")
		{
			streamers.GET("", routes.GetStreamers)
			streamers.GET(":id", routes.GetStreamer)
			streamers.POST("", routes.PostStreamer)
			streamers.PUT(":id", routes.PutStreamer)
			streamers.DELETE(":id", routes.DeleteStreamer)
		}
	}

	router.Run(":3001")
}
