package infrastructure

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // for access to mysql
	"github.com/uma-arai/sbcntr-backend/utils"
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

	DBMS := c.MySQL.DBMS
	USER := c.MySQL.Username
	PASS := c.MySQL.Password
	PROTOCOL := c.MySQL.Protocol
	DBNAME := c.MySQL.DBName

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true&loc=Asia%2FTokyo"
	conn, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		// panic(err.Error())

		fmt.Println("Error: No database connection established.", CONNECT)
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
	return handler.Conn.Model(model).Where(query, args).Count(out)
}

// Create ...
func (handler *SQLHandler) Create(input interface{}) (result interface{}, err error) {
	fmt.Println(input)
	response := handler.Conn.Create(input)

	return response.Value, response.Error
}

// Update ...
func (handler *SQLHandler) Update(out interface{}, value interface{}, query interface{}, args ...interface{}) interface{} {
	return handler.Conn.Model(out).Where(query, args).Updates(value)
}
