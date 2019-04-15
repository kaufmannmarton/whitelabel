package models

type Artist struct {
	Name        string    `json:"name"`
	HeaderImage string    `json:"header-image"`
	Description string    `json:"description"`
	Keywords    string    `json:"keywords"`
	Images      *[]string `json:"images"`

	Instagram *string `json:"instagram"`
	Reddit    *string `json:"reddit"`
	Facebook  *string `json:"facebook"`
	Gfycat    *string `json:"gfycat"`
	YouTube   *string `json:"youtube"`

	Twitter              *string `json:"twitter"`
	TwitterPornhubWidget *string `json:"twitter-pornhub-widget"`

	Fancentro  *string `json:"fancentro"`
	OnlyFans   *string `json:"onlyfans"`
	RedTube    *string `json:"redtube"`
	Modelhub   *string `json:"modelhub"`
	XHamster   *string `json:"xhamster"`
	XVideos    *string `json:"xvideos"`
	Clips4Sale *string `json:"clips4sale"`
	YouPorn    *string `json:"youporn"`

	Pornhub       *string `json:"pornhub"`
	PornhubID     *string `json:"pornhub-id"`
	PornhubTag    *string `json:"pornhub-tag"`
	PornhubVideos []Video

	ManyVids       *string `json:"manyvids"`
	ManyVidsVideos []Video
}
