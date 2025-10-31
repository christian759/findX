package model

import (
	"sync"
)

type NodeID string
type RelationshipType string

type Relationship struct {
	From   NodeID
	To     NodeID
	Type   RelationshipType
	Weight int // optional: how strong or frequent the connection is
}

type SocialGraph struct {
	Nodes map[NodeID]*User
	Edges map[NodeID]map[NodeID]*Relationship
	mu    sync.RWMutex
}
