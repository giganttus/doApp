package helpers

import (
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func UploadFile(ctx *gin.Context, retaurantID int, currentDate string) (string, error) {
	bucket := os.Getenv("BUCKET_NAME")

	stClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return "", err
	}

	f, _, err := ctx.Request.FormFile("file")
	if err != nil {
		return "", err
	}

	defer f.Close()

	filename := fmt.Sprintf("%v_%v", retaurantID, currentDate)

	sw := stClient.Bucket(bucket).Object(filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		return "", err
	}

	if err := sw.Close(); err != nil {
		return "", err
	}

	link := fmt.Sprintf("https://storage.googleapis.com/%v/%v", bucket, filename)
	if err != nil {
		return "", err
	}

	return link, nil
}

func DeleteFile(ctx *gin.Context, retaurantID int, currentDate string) error {
	bucket := os.Getenv("BUCKET_NAME")

	stClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%v_%v", retaurantID, currentDate)

	err = stClient.Bucket(bucket).Object(filename).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}
