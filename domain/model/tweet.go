package model

import (
	"time"
)

type Tweet struct {
	CreatedAt            time.Time         `firestore:"createdAt"`
	IDStr                string            `firestore:"idStr"`
	Text                 string            `firestore:"text"`
	Source               string            `firestore:"source"`
	Truncated            bool              `firestore:"truncated"`
	InReplyToStatusIDStr *string           `firestore:"inReplyToStatusIdStr"`
	InReplyToUserIDStr   *string           `firestore:"inReplyToUserIdStr"`
	InReplyToScreenName  *string           `firestore:"inReplyToScreenName"`
	User                 User              `firestore:"user"`
	Coordinates          *Coordinates      `firestore:"coordinates"`
	Place                *Place            `firestore:"place"`
	QuotedStatusIDStr    *string           `firestore:"quotedStatusIdStr"`
	IsQuoteStatus        bool              `firestore:"isQuoteStatus"`
	QuotedStatus         *Tweet            `firestore:"quotedStatus"`
	RetweetedStatus      *Tweet            `firestore:"retweetedStatus"`
	ReplyCount           int               `firestore:"replyCount"`
	RetweetCount         int               `firestore:"retweetCount"`
	FavoriteCount        *int              `firestore:"favoriteCount"`
	Entities             Entities          `firestore:"entities"`
	ExtendedEntities     *ExtendedEntities `firestore:"extendedEntities"`
	Favorited            *bool             `firestore:"favorited"`
	Retweeted            bool              `firestore:"retweeted"`
	Lang                 *string           `firestore:"lang"`
}

func (s *Tweet) IsPostData() {}
