package graph

import (
	"errors"
	"findX/model"
	"sync"
)

const (
	Followers model.RelationshipType = "followers"
	Following model.RelationshipType = "following"
	Mentions  model.RelationshipType = "mentions"
	Likes     model.RelationshipType = "likes"
	Comments  model.RelationshipType = "comments"
	Tagged    model.RelationshipType = "tagged"
)

type SocialGraph struct {
	Nodes map[model.NodeID]*model.User
	Edges map[model.NodeID]map[model.NodeID]*model.Relationship
	mu    sync.RWMutex
}

func NewSocialGraph() SocialGraph {
	return SocialGraph{
		Nodes: make(map[model.NodeID]*model.User),
		Edges: make(map[model.NodeID]map[model.NodeID]*model.Relationship),
	}
}

func (g *SocialGraph) AddUser(user *model.User) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.Nodes[user.ID] = user
}

func (g *SocialGraph) AddRelationship(from, to model.NodeID, relType *model.RelationshipType, weight int) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if _, ok := g.Nodes[from]; !ok {
		return errors.New("source node not found")
	}
	if _, ok := g.Nodes[to]; !ok {
		return errors.New("target node not found")
	}
	rel := &model.Relationship{From: from, To: to, Type: relType, Weight: weight}
	g.Edges[from] = append(g.Edges[from], rel)
	return nil
}
