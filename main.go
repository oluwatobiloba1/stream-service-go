package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/oluwatobiloba1/stream-service-go/controllers"
)

func main() {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"localhost"})

	videoController := controllers.NewVideoController()
	apiRouter := router.Group("/api")
	videoRouter := apiRouter.Group("/videos")
	videoRouter.POST("/", videoController.StartRecording)
	videoRouter.PUT("/:id", videoController.UploadOctetStreamChunk)
	videoRouter.PUT("/:id/blob", videoController.UploadBlobChunk)
	videoRouter.PUT("/:id/finish-recording", videoController.FinishRecording)
	videoRouter.GET("/", videoController.GetAllVideos)
	videoRouter.GET("/:id", videoController.GetSingleVideo)
	videoRouter.DELETE("/:id", videoController.DeleteVideo)

	router.StaticFile("/", "./index.html")

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
