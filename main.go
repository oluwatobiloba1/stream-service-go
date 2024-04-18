package main

import (
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/oluwatobiloba1/stream-service-go/controllers"
)

// type VideoController struct{}
// func NewVideoController() *VideoController {
// 	return &VideoController{}
// }

func main() {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// videoController := NewVideoController()
	// router.POST("/", videoController.startRecording)
	// router.PUT("/:id", videoController.uploadOctetStreamChunk)
	// router.PUT("/:id/blob", videoController.uploadBlobChunk)
	// router.PUT("/:id/finish-recording", videoController.finishRecording)
	// router.GET("/", videoController.getAllVideos)
	// router.GET("/:id", videoController.getSingleVideo)
	// router.DELETE("/:id", videoController.deleteVideo)

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

// func (vc *VideoController) startRecording(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, "startRecording")
// }
// func (vc *VideoController) uploadOctetStreamChunk(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, "uploadOctetStreamChunk")
// }
// func (vc *VideoController) uploadBlobChunk(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, "uploadBlobChunk")
// }
// func (vc *VideoController) finishRecording(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, "finishRecording")
// }
// func (vc *VideoController) getAllVideos(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, "getAllVideos")
// }
// func (vc *VideoController) getSingleVideo(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, "getSingleVideo")
// }
// func (vc *VideoController) deleteVideo(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, "deleteVideo")
// }
