package model

type NodeID string
type RelationshipType string

type Relationship struct {
	From   NodeID
	To     NodeID
	Type   RelationshipType
	Weight int // optional: how strong or frequent the connection is
}
