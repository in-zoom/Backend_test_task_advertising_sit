package handlers

import (
	"Backend_task_advertising_site/DB"
	"Backend_task_advertising_site/data"
	"Backend_task_advertising_site/upload"
	"Backend_task_advertising_site/validation"
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
    "github.com/julienschmidt/httprouter"
)

type errMessage struct {
	Message string `json:"message"`
}

func AddNewAd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	contentType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || !strings.HasPrefix(contentType, "multipart/") {
		ResponseError(w, 400, err)
	}

	multipartReader := multipart.NewReader(r.Body, params["boundary"])

	arrayOfLinks := make([]string, 0)
	for {
		part, err := multipartReader.NextPart()
		if err == io.EOF {
			break
		}

		if err != nil {
			http.Error(w, "unexpected error when retrieving a part of the message", http.StatusInternalServerError)
			return
		}
		defer part.Close()

		switch part.Header.Get("Content-Type") {
		case "image/jpeg":
			nameLinks, err := upload.UploadPhoto(part)
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

			err = DB.AddNewAd(resultDescription, resultTitle, resultPrice, arrayOfLinks)
			if err != nil {
				ResponseError(w, 500, err)
				return
			}
		}
	}
}

func ResponseError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	errMessage := errMessage{err.Error()}
	json.NewEncoder(w).Encode(errMessage)
}
