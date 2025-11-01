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
	nodes map[model.NodeID]*model.User
	edges map[model.NodeID]map[model.NodeID]*model.Relationship
	mu    sync.RWMutex
}

// NewSocialGraph initializes an empty graph
func NewSocialGraph() *SocialGraph {
	return &SocialGraph{
		nodes: make(map[model.NodeID]*model.User),
		edges: make(map[model.NodeID]map[model.NodeID]*model.Relationship),
	}
}

// AddUser inserts or updates a user node
func (g *SocialGraph) AddUser(user *model.User) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.nodes[user.ID] = user
}

// AddRelationship creates a directed relationship between users
func (g *SocialGraph) AddRelationship(from, to model.NodeID, relType model.RelationshipType, weight int) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if _, ok := g.nodes[from]; !ok {
		return errors.New("source node not found")
	}
	if _, ok := g.nodes[to]; !ok {
		return errors.New("target node not found")
	}

	if _, ok := g.edges[from]; !ok {
		g.edges[from] = make(map[model.NodeID]*model.Relationship)
	}

	g.edges[from][to] = &model.Relationship{From: from, To: to, Type: relType, Weight: weight}
	return nil
}

// GetUser returns a user by ID
func (g *SocialGraph) GetUser(id model.NodeID) (*model.User, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	user, ok := g.nodes[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// GetRelationships returns all relationships for a given user
func (g *SocialGraph) GetRelationships(id model.NodeID) ([]*model.Relationship, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if _, ok := g.nodes[id]; !ok {
		return nil, errors.New("user not found")
	}

	relationships := make([]*model.Relationship, 0, len(g.edges[id]))
	for _, rel := range g.edges[id] {
		relationships = append(relationships, rel)
	}
	return relationships, nil
}

// Followers returns all users that follow a given user
func (g *SocialGraph) GetFollowers(target model.NodeID) []*model.User {
	g.mu.RLock()
	defer g.mu.RUnlock()
	var followers []*model.User
	for from, rels := range g.edges {
		for _, rel := range rels {
			if rel.Type == Followers && rel.To == target {
				if u, ok := g.nodes[from]; ok {
					followers = append(followers, u)
				}
			}
		}
	}
	return followers
}
