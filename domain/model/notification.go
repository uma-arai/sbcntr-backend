package model

type (
	// Notification ... entity for notification db result
	Notification struct {
		ID        int    `json:"id" db:"id"`
		UserId    string `json:"user_id" db:"user_id"`
		Title     string `json:"title" db:"title"`
		Message   string `json:"message" db:"message"`
		IsRead    bool   `json:"is_read" db:"is_read"`
		Type      string `json:"type" db:"type"`
		CreatedAt string `json:"created_at" db:"created_at"`
		UpdatedAt string `json:"updated_at" db:"updated_at"`
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

// TableName ... override table name accessor
func (Notification) TableName() string {
	return "notifications"
}
