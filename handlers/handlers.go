package handlers

import (
	"Backend_task_advertising_site/DB"
	"Backend_task_advertising_site/data"
	"Backend_task_advertising_site/upload"
	"Backend_task_advertising_site/validation"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type errMessage struct {
	Message string `json:"message"`
}

func Upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    multipart, err := r.MultipartReader()
	if err != nil {
		ResponseError(w, 400, err)
	}
	arrayOfLinks, err := upload.UploadPhoto(multipart)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}
	fmt.Println(arrayOfLinks)

}

func AddNewAd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	addedAd := data.NewAd{}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewDecoder(r.Body).Decode(&addedAd)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}

	resultDescription, err := validation.ValidateDescription(addedAd.Description)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}

	resultTitle, err := validation.ValidateTitle(addedAd.Title)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}

	resultPrice, err := validation.ValidatePrice(addedAd.Price)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}

	err = DB.AddNewAd(resultDescription, resultTitle, resultPrice)
	if err != nil {
		ResponseError(w, 500, err)
		return
	}

}

func ResponseError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	errMessage := errMessage{err.Error()}
	json.NewEncoder(w).Encode(errMessage)
}
