package main

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

const (
	S3_REGION = ""
	S3_BUCKET = ""
)

func listFiles() []*s3.Object {
	sess, sessionErr := session.NewSession(&aws.Config{
		Region: aws.String(S3_REGION)},
	)

	if sessionErr != nil {
		fmt.Println("Error Creating aws session: ", sessionErr)
	}

	svc := s3.New(sess)
	fileList, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(S3_BUCKET)})
	if err != nil {
		fmt.Println("Unable to list items in bucket %s, %v", S3_BUCKET, err)
	}

	return fileList.Contents
}

func upload(filename string, reader io.Reader) {
	sess, sessionErr := session.NewSession(&aws.Config{
		Region: aws.String(S3_REGION)},
	)
	svc := s3manager.NewUploader(sess)

	if sessionErr != nil {
		fmt.Println("Error Creating aws session: ", sessionErr)
	}

	fmt.Println("Uploading file to S3...")
	result, uploadError := svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(filepath.Base(filename)),
		Body:   reader,
	})

	if uploadError != nil {
		fmt.Println("Upload Error: ", uploadError)
	}

	fmt.Printf("Successfully uploaded %s to %s\n", filename, result.Location)
}

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.LoadHTMLGlob("html/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", map[string]interface{}{
			"files": listFiles(),
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
