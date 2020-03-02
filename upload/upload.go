package upload

import (
	"Backend_task_advertising_site/random"
	"Backend_task_advertising_site/validation"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"
)

func UploadPhoto(multipart *multipart.Reader) ([]string, error) {
	const IMAGE_DIR = "./image"
	const STAT_IMAGE_PATH = "/stat-img"
	var path, fullFileName, nameLinks string
	arrayOfLinks := make([]string, 0)

	for {
		part, err := multipart.NextPart()
		if err != nil {
			break
		}

		fileName := part.FileName()
		var imgExt string
		arr := strings.Split(fileName, ".")
		if len(arr) > 1 {
			imgExt = arr[len(arr)-1]
		}

		err = validation.ValidateFormatPhoto(imgExt)
		if err != nil {
			return nil, err
		}

		err = os.MkdirAll(fmt.Sprintf("%s/product/%s", IMAGE_DIR, fileName), os.ModePerm)
		if err != nil {
			return nil, err
		}

		path = fmt.Sprintf("%s/product/%s", IMAGE_DIR, fileName)

		fileBytes, err := ioutil.ReadAll(part)
		if err != nil {
			return nil, err
		}

	    fullFileName = fmt.Sprintf("%s.%s", random.RandomFileName(), imgExt)
		fileOnDisk, err := os.Create(fmt.Sprintf("%s/%s", path, fullFileName))
		if err != nil {
			return nil, err
		}
		_, err = fileOnDisk.Write(fileBytes)
		if err != nil {
			return nil, err
		}
		nameLinks = fmt.Sprintf("%s/%s", strings.Replace(path, IMAGE_DIR, STAT_IMAGE_PATH, 1), fullFileName)
		arrayOfLinks = append(arrayOfLinks, nameLinks)
	}
	err := validation.TheNumberOfLinksToThePhoto(arrayOfLinks)
	if err != nil {
		return nil, err
	}
	return arrayOfLinks, nil
}
