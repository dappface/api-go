package firestore

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/dappface/api-go/domain/model"
	"github.com/dappface/api-go/domain/repository"
	"github.com/mitchellh/mapstructure"
)

func NewPostRepository() repository.PostRepository {
	return &PostRepository{}
}

type PostRepository struct{}

func (r *PostRepository) getByID(ctx context.Context, id string) (*model.Post, error) {
	doc, err := Client.
		Collection(Collection.Posts).
		Doc(id).
		Get(ctx)
	if err != nil {
		log.Printf("Failed to get post: %+v", err)
		return nil, err
	}
	var p model.Post
	if err := doc.DataTo(&p); err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PostRepository) GetAll(ctx context.Context, first *int, last *int, before *string, after *string) (*[]model.Post, error) {
	posts := make([]model.Post, 0)
	q := Client.
		Collection(Collection.Posts).
		OrderBy("publishedAt", firestore.Desc)

	if before != nil {
		p, err := r.getByID(ctx, *before)
		if err != nil {
			log.Printf("Failed to get before post: %+v", err)
			return nil, err
		}
		q = q.EndBefore(p.PublishedAt)
	}

	if after != nil {
		p, err := r.getByID(ctx, *after)
		if err != nil {
			log.Printf("Failed to get after post: %+v", err)
			return nil, err
		}
		q = q.StartAfter(p.PublishedAt)
	}

	if last != nil {
		q = q.Limit(*last)
	}

	if first != nil {
		q = q.Limit(*first)
	}

	docs, err := q.Documents(ctx).GetAll()

	if err != nil {
		log.Printf("Error on getting posts: %v", err)
		return nil, err
	}

	for _, doc := range docs {
		var p model.Post
		if err := doc.DataTo(&p); err != nil {
			return nil, err
		}

		p.ID = doc.Ref.ID

		d, err := doc.DataAt("postData")
		if err != nil {
			return nil, err
		}

		switch p.PostType {
		case model.PostType.RSSEntry:
			var e model.RSSEntry
			if err := mapstructure.Decode(d, &e); err != nil {
				return nil, err
			}

			p.PostData = &e
		case model.PostType.Tweet:
			var t model.Tweet
			if err := mapstructure.Decode(d, &t); err != nil {
				return nil, err
			}

			p.PostData = &t
		default:
			return nil, fmt.Errorf("Invalid post type")
		}

		posts = append(posts, p)
	}
	return &posts, nil
}

func (r *PostRepository) HasPage(ctx context.Context, cursor string, pageType string) (*bool, error) {
	q := Client.
		Collection(Collection.Posts).
		OrderBy("publishedAt", firestore.Desc).
		Limit(1)
	p, err := r.getByID(ctx, cursor)
	if err != nil {
		log.Printf("Failed to get HasNextPage: %+v", err)
		return nil, err
	}

	switch pageType {
	case "next":
		q = q.StartAfter(p.PublishedAt)
	case "prev":
		q = q.EndBefore(p.PublishedAt)
	}

	posts, err := q.
		Documents(ctx).
		GetAll()

	if err != nil {
		log.Printf("Failed to get HasNextPage: %+v", err)
		return nil, err
	}

	b := len(posts) != 0

	return &b, err
}
