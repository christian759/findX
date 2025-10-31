package graph

type NodeID string
type RelationshipType string

const (
	Followers RelationshipType = "followers"
	Following RelationshipType = "following"
	Mentions  RelationshipType = "mentions"
	Likes     RelationshipType = "likes"
	Comments  RelationshipType = "comments"
	Tagged    RelationshipType = "tagged"
)
