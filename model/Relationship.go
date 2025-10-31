package model

import (
	"findX/graph"
)

type Relationship struct {
	From   graph.NodeID
	To     graph.NodeID
	Type   graph.RelationshipType
	Weight int // optional: how strong or frequent the connection is
}
