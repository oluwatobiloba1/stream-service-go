package controllers

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	// "github.com/lucsky/cuid"
)

type VideoController struct{}

// type StructA struct {
// 	chunk `form:"chunk"`
// }

func NewVideoController() *VideoController {
	return &VideoController{}
}

func (vc *VideoController) StartRecording(context *gin.Context) {
	// id := cuid.New()
	id := "clv78jlrk0000v8vgp9bw0kb6"
	filename := id + ".mp4"
	folder := "uploads/"
	file, err := os.Create(folder + filename)
	if err != nil {
		log.Fatal(err)
		context.JSON(http.StatusInternalServerError, "Error creating the file")
		return
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
	context.IndentedJSON(http.StatusOK, id)
}
func (vc *VideoController) UploadOctetStreamChunk(context *gin.Context) {
	id, exists := context.Params.Get("id")
	if !exists {
		context.JSON(http.StatusInternalServerError, "ID param must be present")
		return
	}
	// // get the chunk from request body
	// bodySize, _ := strconv.ParseInt(context.Request.Header.Get("Content-Length"), 10, 64)
	// buff := bytes.NewBuffer(make([]byte, 0, bodySize))
	// _, err := io.Copy(buff, context.Request.Body)
	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, "Error reading chunk from request")
	// 	return
	// }
	// fmt.Println("buff: ", buff)

	// // check for filename matching the ID
	// // open the file
	// // append the new chunk to the file
	// // retuen success response
	// context.IndentedJSON(http.StatusOK, "uploadOctetStreamChunk")

	// Read request chunk into a buffer
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	// Decode base64 data
	chunk := strings.Split(string(body), ",")[1]
	data, err := base64.StdEncoding.DecodeString(chunk)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode base64 data"})
		return
	}

	// Open or create a file for appending
	filePath := filepath.Join("uploads", id+".mp4")
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file for appending"})
		return
	}
	defer file.Close()

	// Append the decoded data to the file
	_, err = file.Write(data)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to append data to file"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Data appended to file successfully"})
}
func (vc *VideoController) UploadBlobChunk(context *gin.Context) {

	fileData, err := context.FormFile("chunk")
	if err != nil {
		context.JSON(http.StatusInternalServerError, "Error getting file from request")
		return
	}
	filePath := filepath.Join("uploads", "clv78jlrk0000v8vgp9bw0kb6"+".mp4")
	// file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file for appending"})
	// 	return
	// }
	// defer file.Close()
	err2 := context.SaveUploadedFile(fileData, filePath)

	// Append the decoded data to the file
	if err2 != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to append data to file"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Data appended to file successfully"})

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
