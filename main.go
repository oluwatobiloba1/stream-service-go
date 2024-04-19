package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oluwatobiloba1/stream-service-go/controllers"
)

func main() {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"localhost"})

	videoController := controllers.NewVideoController()
	router.POST("/", videoController.StartRecording)
	router.PUT("/:id", videoController.UploadOctetStreamChunk)
	router.PUT("/:id/blob", videoController.UploadBlobChunk)
	router.PUT("/:id/finish-recording", videoController.FinishRecording)
	router.GET("/", videoController.GetAllVideos)
	router.GET("/:id", videoController.GetSingleVideo)
	router.DELETE("/:id", videoController.DeleteVideo)

	router.Run("localhost:8080")
}
