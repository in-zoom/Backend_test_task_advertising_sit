package controllers

import (
	"backend_task_advertising_site/random"
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type BucketController struct {
	Uploader *s3manager.Uploader
}

func (bc BucketController) UploadFilesToBucket(part *multipart.Part) (string, error) {
	fileBytes, err := ioutil.ReadAll(part)
	if err != nil {
		return "", err
	}

	file := io.MultiReader(bytes.NewReader(fileBytes))
	fileName, err := random.RandomFileName(part.FileName())
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
