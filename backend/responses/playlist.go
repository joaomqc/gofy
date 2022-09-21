package responses

type SimplePlaylist struct {
	Id         string      `json:"id"`
	Title      string      `json:"title"`
	Thumbnails []Thumbnail `json:"thumbnails"`
}
