package model

import (
	"time"
)

type Post struct {
	ID        int
	Platform  string
	AuthorID  string
	Text      string
	CreatedAt time.Time
	URLs      []string
	Mentions  []string
	Hashtags  []string
	MediaIDs  []string
	// add association with user
	User []*User
}
