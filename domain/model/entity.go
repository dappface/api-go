package model

import "time"

type Entities struct {
	Hashtags     []Hashtag     `firestore:"hashtags"`
	Media        []Media       `firestore:"media"`
	Symbols      []Symbol      `firestore:"symbols"`
	URLs         []URL         `firestore:"urls"`
	UserMentions []UserMention `firestore:"userMentions"`
	Polls        []Poll        `firestore:"polls"`
}

type ExtendedEntities struct {
	Media []Media `firestore:"media"`
}

type Hashtag struct {
	Indices []int  `firestore:"indices"`
	Text    string `firestore:"text"`
}

type URL struct {
	DisplayURL  string   `firestore:"displayUrl"`
	ExpandedURL string   `firestore:"expandedUrl"`
	Indices     []int    `firestore:"indices"`
	URL         string   `firestore:"url"`
	Unwound     *Unwound `firestore:"unwound"`
}

type Unwound struct {
	URL         string `firestore:"url"`
	Status      int    `firestore:"status"`
	Title       string `firestore:"title"`
	Description string `firestore:"description"`
}

type UserMention struct {
	IDStr      string `firestore:"idStr"`
	Indices    []int  `firestore:"indices"`
	Name       string `firestore:"name"`
	ScreenName string `firestore:"screenName"`
}

type Symbol struct {
	Indices []int  `firestore:"indices"`
	Text    string `firestore:"text"`
}

type Poll struct {
	Options     []Option  `firestore:"options"`
	EndDatetime time.Time `firestore:"endDatetime"`
}

type Option struct {
	Position int    `firestore:"position"`
	Text     string `firestore:"text"`
}
