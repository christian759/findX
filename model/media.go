package model

type Media struct {
	ID        string
	PostID    string
	URL       string
	Mime      string
	Size      int64
	Hash      string // phash/dhash
	EXIF      map[string]string
	LocalPath string
}
