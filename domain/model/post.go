package model

import "time"

var (
	PostType = postType{
		RSSEntry: "rssEntry",
		Tweet:    "tweet",
	}
)

type Post struct {
	ID          string    `firestore:"-"`
	PostDataID  string    `firestore:"postDataId"`
	PostType    string    `firestore:"postType"` // "rssEntry" | "tweet"
	PostData    PostData  `firestore:"-"`
	Language    *string   `firestore:"language"`
	PublishedAt time.Time `firestore:"publishedAt"`
	CreatedAt   time.Time `firestore:"createdAt"`
	UpdatedAt   time.Time `firestore:"updatedAt"`
}

func (p *Post) IsNode() {}

type PostData interface {
	IsPostData()
}

type postType struct {
	RSSEntry string
	Tweet    string
}
