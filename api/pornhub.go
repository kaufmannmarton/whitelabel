package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"whitelabel/models"
)

const pornhubBaseURL = "http://www.pornhub.com/webmasters/search"

func GetLatestPornhubVideos(pornhubID string) ([]models.Video, error) {
	return getPornhubVideos(pornhubBaseURL + "?ordering=newest&phrase[]=" + pornhubID)
}

func GetMostViewedPornhubVideos(pornhubID string) ([]models.Video, error) {
	return getPornhubVideos(pornhubBaseURL + "?ordering=mostviewed&phrase[]=" + pornhubID)
}

func getPornhubVideos(url string) ([]models.Video, error) {
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

	return data["videos"], nil
}
