package main

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

type fileMetadata struct {
	Filename string
}

const (
	S3_REGION = ""
	S3_BUCKET = ""
)

func getFiles() []fileMetadata {
	files := make([]fileMetadata, 0)

	files = append(files, fileMetadata{
		Filename: "test.txt",
	})

	return files
}

// func download()

func upload(filename string, reader io.Reader) {
	conf := aws.Config{Region: aws.String(S3_REGION)}
	sess := session.New(&conf)
	svc := s3manager.NewUploader(sess)

	fmt.Println("Uploading file to S3...")
	result, err := svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(filepath.Base(filename)),
		Body:   reader,
	})
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Printf("Successfully uploaded %s to %s\n", filename, result.Location)
}

func main() {
	session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.LoadHTMLGlob("html/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", map[string]interface{}{
			"files": getFiles(),
		})
	})

	router.GET("/download/:filename", func(ctx *gin.Context) {
		uploadsPath := "./uploads/"
		fileName := ctx.Param("filename")
		// If this was a real application this should have more security and validation
		targetPath := filepath.Join(uploadsPath, fileName)
		fmt.Printf("file: %s", targetPath)
		ctx.Header("Content-Description", "File Transfer")
		ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.Header("Content-Disposition", "attachment; filename="+fileName)
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.File(targetPath)
	})

	router.POST("/upload", func(ctx *gin.Context) {
		// Multipart form
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		files := form.File["files"]

		for _, fileHeader := range files {
			reader, err := fileHeader.Open()
			if err != nil {
				ctx.String(http.StatusInternalServerError, fmt.Sprintf("get form err: %s", err.Error()))
				return
			}
			upload(fileHeader.Filename, reader)
		}

		ctx.Redirect(http.StatusFound, "/")
	})
	router.Run("localhost:8080")
}
