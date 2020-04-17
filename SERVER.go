package main

import (
	"backend_task_advertising_site/controllers"
	"backend_task_advertising_site/handlers"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	initialize()
	ec := handlers.EventsController{BucketCtrl: initBucketController()}
	router := httprouter.New()
	router.POST("/create", ec.AddNewAd)
	router.GET("/ads", handlers.GetListAds)
	router.GET("/ad", handlers.GetSpecificAd)
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func initialize() {
	if err := godotenv.Load(); err != nil {
		log.Print("Файл .env не найден")
	}
}

func initBucketController() *controllers.BucketController {
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
	uploader := s3manager.NewUploader(s)
	return &controllers.BucketController{Uploader: uploader}
}
