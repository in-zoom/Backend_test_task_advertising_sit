package controllers

import (
	"backend_task_advertising_site/random"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type BucketController struct {
	Uploader *s3manager.Uploader
}

func (bc BucketController) UploadFilesToBucket(part *multipart.Part) (string, error) {

	nameFile := part.FileName()
	file, err := os.Open(nameFile)
	if err != nil {
		return "", err
	}

	fileName, err := random.RandomFileName(nameFile)
	if err != nil {
		return "", err
	}

	result, err := bc.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("pictures-for-links"),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		return "", err
	}
	return result.Location, nil
}
