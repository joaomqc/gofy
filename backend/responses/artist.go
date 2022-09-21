package responses

type SimpleArtist struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	Thumbnails []Thumbnail `json:"thumbnails"`
}
