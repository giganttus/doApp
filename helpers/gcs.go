package helpers

import (
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func UploadFile(ctx *gin.Context) (string, error) {
	bucket := os.Getenv("BUCKET_NAME")

	stClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return "", err
	}

	f, uploadedFile, err := ctx.Request.FormFile("file")
	if err != nil {
		return "", err
	}

	defer f.Close()

	sw := stClient.Bucket(bucket).Object(uploadedFile.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		return "", err
	}

	if err := sw.Close(); err != nil {
		return "", err
	}

	link := fmt.Sprintf("https://storage.googleapis.com/%v/%v", bucket, uploadedFile.Filename)
	if err != nil {
		return "", err
	}
	return link, nil
}
