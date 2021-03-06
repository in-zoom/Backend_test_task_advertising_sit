package handlers

import (
	"backend_task_advertising_site/DB"
	"backend_task_advertising_site/controllers"
	"backend_task_advertising_site/data"
	"backend_task_advertising_site/validation"
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type EventsController struct {
	BucketCtrl *controllers.BucketController
}

type errMessage struct {
	Message string `json:"message"`
}

type message struct {
	OkMessage string `json:"okMessage"`
	Status    int    `json:"status"`
	Id        int    `json:"id"`
}

func (ec EventsController) AddNewAd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	contentType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || !strings.HasPrefix(contentType, "multipart/") {
		ResponseError(w, 400, err)
	}

	multipartReader := multipart.NewReader(r.Body, params["boundary"])
	var resultDescription, resultTitle string
	var resultPrice float64
	arrayOfLinks := make([]string, 0)
	for {
		part, err := multipartReader.NextPart()
		if err == io.EOF {
			break
		}

		if err != nil {
			ResponseError(w, 500, err)
			return
		}
		defer part.Close()

		switch part.Header.Get("Content-Type") {
		case "image/jpeg":

			nameLinks, err := ec.BucketCtrl.UploadFilesToBucket(part)
			if err != nil {
				ResponseError(w, 400, err)
				return
			}

			arrayOfLinks = append(arrayOfLinks, nameLinks)
			err = validation.TheNumberOfLinksToThePhoto(arrayOfLinks)
			if err != nil {
				ResponseError(w, 400, err)
				return
			}

		case "application/json":
			addedAd := data.NewAd{}
			err = json.NewDecoder(part).Decode(&addedAd)
			if err != nil {
				ResponseError(w, 400, err)
				return
			}

			resultDescription, err = validation.ValidateDescription(addedAd.Description)
			if err != nil {
				ResponseError(w, 400, err)
				return
			}

			resultTitle, err = validation.ValidateTitle(addedAd.Title)
			if err != nil {
				ResponseError(w, 400, err)
				return
			}

			resultPrice, err = validation.ValidatePrice(addedAd.Price)
			if err != nil {
				ResponseError(w, 400, err)
				return
			}
		}
	}

	id, err := DB.AddNewAd(resultDescription, resultTitle, resultPrice, arrayOfLinks)
	if err != nil {
		ResponseError(w, 500, err)
		return
	}
	ResponceOk(w, 200, id)
}

func GetListAds(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	url := r.URL.Query()
	attribute := url.Get("atribute")
	order := url.Get("order")
	offset := url.Get("offset")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	resultAtribute, err := validation.ValidateAtribute(attribute)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}

	resultOrder, err := validation.ValidateOrder(order)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}

	resultOffset, err := validation.ValidateOffset(offset, DB.Сonnect())
	if err != nil {
		ResponseError(w, 400, err)
		return
	}

	listAds, err := DB.ReceiveListAds(resultAtribute, resultOrder, resultOffset)
	if err != nil {
		ResponseError(w, 500, err)
		return
	}
	json.NewEncoder(w).Encode(listAds)
}

func GetSpecificAd(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	url := r.URL.Query()
	id := url.Get("id")
	fields, _ := url["fields"]

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//db := DB.Сonnect()
	resultId, err := validation.ValidateId(id, DB.Сonnect())
	if err != nil {
		ResponseError(w, 400, err)
		return
	}

	resultfields, err := validation.ValidateFields(fields)
	if err != nil {
		ResponseError(w, 400, err)
		return
	}

	ad, err := DB.GetOneAd(resultId, resultfields)
	if err != nil {
		ResponseError(w, 500, err)
		return
	}
	json.NewEncoder(w).Encode(ad)
}

func ResponceOk(w http.ResponseWriter, code int, id int) {
	w.WriteHeader(code)
	m := message{"Ваше объявление успешно добавленно", code, id}
	json.NewEncoder(w).Encode(m)
}

func ResponseError(w http.ResponseWriter, code int, err error) {
	errMessage := errMessage{err.Error()}
	json.NewEncoder(w).Encode(errMessage)
}
