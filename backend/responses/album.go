package responses

type SimpleAlbum struct {
	Id         string      `json:"id"`
	Title      string      `json:"title"`
	Thumbnails []Thumbnail `json:"thumbnails"`
}
