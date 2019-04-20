package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"whitelabel/models"
)

const pornhubBaseURL = "http://www.pornhub.com/webmasters"

func GetLatestPornhubVideos(id string) ([]models.Video, error) {
	return getPornhubVideos(pornhubBaseURL + "/search?ordering=newest&phrase[]=" + id)
}

func GetLatestPornhubVideosByTag(tag string) ([]models.Video, error) {
	return getPornhubVideos(pornhubBaseURL + "/search?ordering=newest&tags[]=" + tag)
}

func GetMostViewedPornhubVideos(id string) ([]models.Video, error) {
	return getPornhubVideos(pornhubBaseURL + "/search?ordering=mostviewed&phrase[]=" + id)
}

func GetMostViewedPornhubVideosByTag(tag string) ([]models.Video, error) {
	return getPornhubVideos(pornhubBaseURL + "/search?mostviewed=newest&tags[]=" + tag)
}

func GetPornhubVideoById(id string) (*models.Video, error) {
	return getPornhubVideo(pornhubBaseURL + "/video_by_id?id=" + id)
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

func getPornhubVideo(url string) (*models.Video, error) {
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

	var data map[string]*models.Video

	json.Unmarshal(body, &data)

	return data["video"], nil
}
