package model

// Favorite ...
type Favorite struct {
	Id     string `json:"id"`
	PetId  string `json:"pet_id"`
	UserId string `json:"user_id"`
	Value  bool   `json:"value"`
}
