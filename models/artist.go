package models

type Artist struct {
	Name        string    `json:"name"`
	HeaderImage string    `json:"header-image"`
	Description string    `json:"description"`
	Images      *[]string `json:"images"`

	Instagram *string `json:"instagram"`
	Reddit    *string `json:"reddit"`
	Twitter   *string `json:"twitter"`
	Facebook  *string `json:"facebook"`
	Gfycat    *string `json:"gfycat"`
	YouTube   *string `json:"youtube"`

	Fancentro *string `json:"fancentro"`
	ManyVids  *string `json:"manyvids"`
	OnlyFans  *string `json:"onlyfans"`
	RedTube   *string `json:"redtube"`
	YouPorn   *string `json:"youporn"`
	Modelhub  *string `json:"modelhub"`
	XHamster  *string `json:"xhamster"`
	XVideos   *string `json:"xvideos"`

	Pornhub       *string `json:"pornhub"`
	PornhubID     *string `json:"pornhub-id"`
	PornhubVideos []Video
}
