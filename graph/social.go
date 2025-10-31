package graph

import (
	"findX/model"
)

const (
	Followers model.RelationshipType = "followers"
	Following model.RelationshipType = "following"
	Mentions  model.RelationshipType = "mentions"
	Likes     model.RelationshipType = "likes"
	Comments  model.RelationshipType = "comments"
	Tagged    model.RelationshipType = "tagged"
)

func NewSocialGraph() *model.SocialGraph {
	return &model.SocialGraph{
		Nodes: make(map[model.NodeID]*model.User),
		Edges: make(map[model.NodeID]map[model.NodeID]*model.Relationship),
	}
}
