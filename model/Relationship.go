package model

type Relationship struct {
	ID       int64
	UserID   int64
	TargetID int64
	Type     RelationshipType
}
