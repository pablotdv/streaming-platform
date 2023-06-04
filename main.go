package main

import (
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

func main() {
	router := gin.Default()
	api := router.Group("/api")
	{
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
