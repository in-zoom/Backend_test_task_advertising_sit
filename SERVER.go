package main

import (
	"Backend_task_advertising_site/handlers"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	initialize()
	router := httprouter.New()
	router.POST("/create", handlers.AddNewAd)
	router.GET("/ads", handlers.GetListAds)
	router.GET("/ad", handlers.GetSpecificAd)
	http.ListenAndServe(":8080", router)
}

func initialize() {
	if err := godotenv.Load(); err != nil {
		log.Print("Файл .env не найден")
	}
}
