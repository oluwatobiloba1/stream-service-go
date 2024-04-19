package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lucsky/cuid"
)

type VideoController struct{}

func NewVideoController() *VideoController {
	return &VideoController{}
}

func (vc *VideoController) StartRecording(context *gin.Context) {
	id := cuid.New()
	filename := id + ".webm"
	folder := "uploads/"
	file, err := os.Create(folder + filename)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, "Error creating the file")
		return
	}
	if err := file.Close(); err != nil {
		fmt.Println(err)
	}
	context.IndentedJSON(http.StatusOK, id)
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
