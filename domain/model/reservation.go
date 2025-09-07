package model

type Reservation struct {
	PetId           string `json:"pet_id"`
	UserId          string `json:"user_id"`
	Email           string `json:"email"`
	FullName        string `json:"full_name"`
	ReservationDate string `json:"reservation_date"`
}
