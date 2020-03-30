package controllers

import (
	"Backend_task_advertising_site/random"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFilesToBucket(part *multipart.Part) (string, error) {

	nameFile := part.FileName()
	file, err := os.Open(nameFile)
	if err != nil {
		return "", err
	}

	fileName, err := random.RandomFileName(nameFile)
	if err != nil {
		return "", err
	}

	s := initBucketController()
	uploader := s3manager.NewUploader(s)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("pictures-for-links"),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		return "", err
	}
	return result.Location, nil
}

func initBucketController() *session.Session {
	var accessKeyID = os.Getenv("ACCESS_KEY_ID")
	var sectetAccessKey = os.Getenv("SECRET_ACCESS_KEY")

	s, _ := session.NewSession(&aws.Config{
		Endpoint: aws.String("storage.yandexcloud.net"),
		Region:   aws.String("ru-central1"),
		Credentials: credentials.NewStaticCredentials(
			accessKeyID,     // id
			sectetAccessKey, // secret
			""),             // token can be left blank for now
	})

	return s
}
