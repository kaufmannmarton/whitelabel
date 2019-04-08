package models

type Artist struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	HeaderImage string    `json:"header-image"`
	Description string    `json:"description"`
	Images      *[]string `json:"images"`

	Instagram *string `json:"instagram"`
	Reddit    *string `json:"reddit"`
	Twitter   *string `json:"twitter"`
	Facebook  *string `json:"facebook"`
	Gfycat    *string `json:"gfycat"`

	Fancentro *string `json:"fancentro"`
	ManyVids  *string `json:"manyvids"`
	OnlyFans  *string `json:"onlyfans"`
	Pornhub   *string `json:"pornhub"`
	RedTube   *string `json:"redtube"`
	YouPorn   *string `json:"youporn"`
	Modelhub  *string `json:"modelhub"`
	XHamster  *string `json:"xhamster"`

	Videos *[]Video
}
