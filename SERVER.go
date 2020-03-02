package main

import (
	"Backend_task_advertising_site/handlers"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func initialize() {
	if err := godotenv.Load(); err != nil {
		log.Print("Файл .env не найден")
	}
}

func main() {
	initialize()
	router := httprouter.New()
	router.POST("/create", handlers.AddNewAd)
    router.POST("/upload", handlers.Upload)
	http.ListenAndServe(":8080", router)
}
