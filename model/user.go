package model

import (
	"time"
)

type User struct {
	Name      string
	Platform  string
	Handle    string
	Bio       string
	Location  string
	Followers int
	Following int
	CreatedAt time.Time
	// add GraphRelation []Relation
}
