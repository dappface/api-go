package model

import (
	"time"
)

type User struct {
	IDStr                string        `firestore:"idStr"`
	Name                 string        `firestore:"name"`
	ScreenName           string        `firestore:"screenName"`
	Location             *string       `firestore:"location"`
	URL                  *string       `firestore:"url"`
	Entities             *UserEntities `firestore:"entities"`
	Description          *string       `firestore:"description"`
	Protected            bool          `firestore:"protected"`
	Verified             bool          `firestore:"verified"`
	FollowersCount       int           `firestore:"followersCount"`
	FriendsCount         int           `firestore:"friendsCount"`
	ListedCount          int           `firestore:"listedCount"`
	FavouritesCount      int           `firestore:"favouritesCount""`
	StatusesCount        int           `firestore:"statusesCount"`
	CreatedAt            time.Time     `firestore:"createdAt"`
	Lang                 string        `firestore:"lang"`
	ProfileImageURLHTTPS string        `firestore:"profileImageUrlHttps"`
}

type UserEntities struct {
	URL UserURL `firestore:"url"`
}

type UserURL struct {
	URLs []URL `firestore:"urls"`
}
