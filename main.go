package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type fileMetadata struct {
	Filename string
}

func getFiles() []fileMetadata {
	files := make([]fileMetadata, 0)

	files = append(files, fileMetadata{
		Filename: "test.txt",
	})

	return files
}

func main() {
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

		for _, file := range files {
			filename := filepath.Join("./uploads", filepath.Base(file.Filename))
			if err := ctx.SaveUploadedFile(file, filename); err != nil {
				ctx.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}

		ctx.Redirect(http.StatusFound, "/")
	})
	router.Run("localhost:8080")
}
