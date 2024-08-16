package entity

type Photos struct {
	ID           int    `json:"ID"`
	AlbumId      int    `json:"albumId"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}
