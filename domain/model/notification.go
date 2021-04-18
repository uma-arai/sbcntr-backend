package model

// Notification ... entity for notification db result
type Notification struct {
	ID          string `json:"id" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Category    string `json:"category" gorm:"column:category"`
	Read        bool   `json:"read" gorm:"column:read"`
	CreatedAt   string `json:"createdAt" gorm:"column:createdAt"`
}
