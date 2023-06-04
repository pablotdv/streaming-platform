package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pablotdv/streaming-platform/data"
	"github.com/pablotdv/streaming-platform/models"
	"github.com/pablotdv/streaming-platform/restreamer"
	"github.com/pablotdv/streaming-platform/schemas"
	"gorm.io/gorm"
)

func GetStreamers(c *gin.Context) {
	var streamers []schemas.StreamerGetResponse
	if err := data.Db.Model(&models.Streamer{}).Select("id, name, url_stream, url_player").Find(&streamers).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, streamers)
}

func GetStreamer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var streamer schemas.StreamerGetResponse
	if err := data.Db.Model(&models.Streamer{}).Where("id = ?", id).Select("id, name, url_stream, url_player").Find(&streamer).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, streamer)
}

func PostStreamer(c *gin.Context) {
	var json schemas.StreamerPostRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	streamer := models.Streamer{
		Name:         json.Name,
		RestreamerId: uuid.New().String(),
	}

	data.Db.Transaction(func(tx *gorm.DB) error {
		process, err := restreamer.CreateProcess(streamer.RestreamerId)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return err
		}
		streamer.ProcessId = process.ID
		streamer.UrlStream = process.Input[0].Address
		streamer.UrlPlayer = "http://localhost:8080/memfs/" + streamer.RestreamerId + ".m3u8"

		if err := tx.Create(&streamer).Error; err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return err
		}
		return nil
	})

	response := schemas.StreamerPostResponse{
		Name:      streamer.Name,
		UrlStream: streamer.UrlStream,
		UrlPlayer: streamer.UrlPlayer,
	}
	c.JSON(http.StatusOK, response)
}

func PutStreamer(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func DeleteStreamer(c *gin.Context) {
	c.JSON(200, gin.H{})
}
