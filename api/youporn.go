package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"whitelabel/models"
)

const youpornBaseURL = "http://www.youporn.com/api/webmasters/search"

func GetMostViewedYouPornVideos(youpornID string) ([]models.Video, error) {
	return getYouPornVideos(youpornBaseURL + "?ordering=mostviewed&search=" + youpornID)
}

func getYouPornVideos(url string) ([]models.Video, error) {
	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)

		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)

		return nil, err
	}

	var data map[string][]models.Video

	json.Unmarshal(body, &data)

	return data["video"], nil
}
