package model

type Media struct {
	DisplayURL        string  `firestore:"displayUrl"`
	ExpandedURL       string  `firestore:"expandedUrl"`
	IDStr             string  `firestore:"idStr"`
	Indices           []int   `firestore:"indices"`
	MediaURL          string  `firestore:"mediaUrl"`
	MediaURLHTTPS     string  `firestore:"mediaUrlHttps"`
	Sizes             Sizes   `firestore:"sizes"`
	SourceStatusIDStr *string `firestore:"sourceStatusIdStr"`
	Type              string  `firestore:"type"`
	URL               string  `firestore:"url"`
}

type Sizes struct {
	Thumb  Size `firestore:"thumb"`
	Large  Size `firestore:"large"`
	Medium Size `firestore:"medium"`
	Small  Size `firestore:"small"`
}

type Size struct {
	W      int    `firestore:"w"`
	H      int    `firestore:"h"`
	Resize string `firestore:"resize"`
}
