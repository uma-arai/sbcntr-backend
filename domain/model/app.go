package model

// App ... entity for db result
type App struct {
	ID        string    `json:"id" gorm:"column:id"`
	Message   string    `json:"message" gorm:"column:message"`
}
