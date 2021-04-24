package model

// App ... entity for db result
type App struct {
	ID      string `json:"id" gorm:"column:id"`
	Message string `json:"message" gorm:"column:message"`
}

// Response ...
type Response struct {
	Code    int    `json:"code" xml:"code"`
	Message string `json:"msg" xml:"msg"`
}

// Hello ... entity for hello message
type Hello struct {
	Data string `json:"data" xml:"data"`
}

type (
	// Item ...
	Item struct {
		ID        int    `json:"id" gorm:"column:id"`
		Title     string `json:"title" gorm:"column:title"`
		Name      string `json:"name" gorm:"column:name"`
		Favorite  bool   `json:"favorite" gorm:"column:favorite"`
		Img       string `json:"img" gorm:"column:img"`
		CreatedAt string `json:"createdAt" gorm:"column:createdAt"`
		UpdatedAt string `json:"updatedAt" gorm:"column:updatedAt"`
	}

	// Items ...
	Items struct {
		Data []Item `json:"data"`
	}
)

// TableName ... override GORM table name accessor
func (Item) TableName() string {
	return "Item"
}
