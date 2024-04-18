package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoController struct{}

func NewVideoController() *VideoController {
	return &VideoController{}
}

func (vc *VideoController) StartRecording(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "startRecording")
}
func (vc *VideoController) UploadOctetStreamChunk(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "uploadOctetStreamChunk")
}
func (vc *VideoController) UploadBlobChunk(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "uploadBlobChunk")
}
func (vc *VideoController) FinishRecording(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "finishRecording")
}
func (vc *VideoController) GetAllVideos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "getAllVideos")
}
func (vc *VideoController) GetSingleVideo(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "getSingleVideo")
}
func (vc *VideoController) DeleteVideo(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "deleteVideo")
}
