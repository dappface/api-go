package model

type Coordinates struct {
	Coordinates     []float64 `firestore:"coordinates"` // [longitude, latitude]
	CoordinatesType string    `firestore:"coordinatesType"`
}

type Place struct {
	ID          string `firestore:"id"`
	URL         string `firestore:"url"`
	PlaceType   string `firestore:"placeType"`
	Name        string `firestore:"name"`
	FullName    string `firestore:"fullName"`
	CountryCode string `firestore:"countryCode"`
	Country     string `firestore:"country"`
}
