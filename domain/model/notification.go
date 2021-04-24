package model

type (
	// Notification ... entity for notification db result
	Notification struct {
		ID          int    `json:"id" gorm:"column:id"`
		Title       string `json:"title" gorm:"column:title"`
		Description string `json:"description" gorm:"column:description"`
		Category    string `json:"category" gorm:"column:category"`
		Unread      bool   `json:"unread" gorm:"column:unread"`
		CreatedAt   string `json:"createdAt" gorm:"column:createdAt"`
		UpdatedAt   string `json:"updatedAt" gorm:"column:updatedAt"`
	}

	// Notifications ... array entity for notification
	Notifications struct {
		Data []Notification `json:"data"`
	}

	// NotificationCount ... unread count of notifications
	NotificationCount struct {
		Data int `json:"data"`
	}
)

// TableName ... override GORM table name accessor
func (Notification) TableName() string {
	return "Notification"
}
