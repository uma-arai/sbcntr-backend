package model

// App ... entity for db result
type App struct {
	ID      string `json:"id" db:"id"`
	Message string `json:"message" db:"message"`
}

type InputUpdateLikeRequest struct {
	PetId  string `json:"pet_id"`
	UserId string `json:"user_id"`
	Value  bool   `json:"value"`
}

// Response ...
type Response struct {
	Code    int    `json:"code" xml:"code"`
	Message string `json:"msg" xml:"msg"`
}

// APIResponse ...
type APIResponse struct {
	Data interface{} `json:"data"`
}

// PetFilter ... フィルタ
type PetFilter struct {
	ID              string  `query:"id"`
	Name            string  `query:"name"`
	Breed           string  `query:"breed"`
	Gender          string  `query:"gender"`
	Price           float64 `query:"price"`
	ReferenceNumber string  `query:"reference_number"`
}
