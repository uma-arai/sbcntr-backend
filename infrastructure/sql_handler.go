package infrastructure

import (
	"fmt"
	"gorm.io/driver/mysql"

	"github.com/uma-arai/sbcntr-backend/utils"
	_ "gorm.io/driver/mysql" // for access to mysql
	"gorm.io/gorm"
)

// DB ...
type DB struct {
	Host     string
	Username string
	Password string
	DBName   string
	Connect  *gorm.DB
}

// SQLHandler ... SQL handler struct
type SQLHandler struct {
	Conn *gorm.DB
}

// NewSQLHandler ...
func NewSQLHandler() *SQLHandler {

	c := utils.NewConfigDB()

	USER := c.MySQL.Username
	PASS := c.MySQL.Password
	PROTOCOL := c.MySQL.Protocol
	DBNAME := c.MySQL.DBName

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true&loc=Asia%2FTokyo"
	conn, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})

	if err != nil {
		// Note: ハンズオンでパニックすると辛いのでエラーハンドリングはコメントアウトしています
		// panic(err.Error())
		fmt.Print("Error: No database connection established.")
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}

// Where ...
func (handler *SQLHandler) Where(out interface{}, query interface{}, args ...interface{}) interface{} {
	if query == "" {
		return handler.Conn.Find(out)
	}

	return handler.Conn.Where(query, args).Find(out)
}

// Scan ...
func (handler SQLHandler) Scan(out interface{}, order string) interface{} {
	return handler.Conn.Order(order).Find(out)
}

// Count ...
func (handler *SQLHandler) Count(out *int, model interface{}, query interface{}, args ...interface{}) interface{} {
	out64 := int64(*out)
	return handler.Conn.Model(model).Where(query, args).Count(&out64)
}

// Create ...
func (handler *SQLHandler) Create(input interface{}) (result interface{}, err error) {
	fmt.Println(input)
	response := handler.Conn.Create(input)

	return nil, response.Error
}

// Update ...
func (handler *SQLHandler) Update(out interface{}, value interface{}, query interface{}, args ...interface{}) interface{} {
	return handler.Conn.Model(out).Where(query, args).Updates(value)
}
