// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Connection interface {
	IsConnection()
}

type Edge interface {
	IsEdge()
}

type Node interface {
	IsNode()
}

type FeedConnection struct {
	Edges    []*PostEdge `json:"edges"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

func (FeedConnection) IsConnection() {}

type PageInfo struct {
	HasNextPage     bool `json:"hasNextPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
}

type PostEdge struct {
	Cursor string `json:"cursor"`
	Node   *Post  `json:"node"`
}

func (PostEdge) IsEdge() {}
