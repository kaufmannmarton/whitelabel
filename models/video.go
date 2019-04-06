package models

type Video struct {
	ID           string     `json:"video_id"`
	Duration     string     `json:"duration"`
	Title        string     `json:"title"`
	URL          string     `json:"url"`
	DefaultThumb string     `json:"default_thumb"`
	Thumb        string     `json:"thumb"`
	Thumbs       []Thumb    `json:"thumbs"`
	Tags         []Tag      `json:"tags"`
	Categories   []Category `json:"categories"`
}
