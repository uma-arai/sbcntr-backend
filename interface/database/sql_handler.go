package database

// SQLHandler ...
type SQLHandler interface {
	Where(out interface{}, query interface{}, args ...interface{}) interface{}
	Scan(out interface{}, order string) interface{}
	Count(out *int, model interface{}, query interface{}, args ...interface{}) interface{}
	Create(input interface{}) (result interface{}, err error)
	Update(out interface{}, value interface{}, query interface{}, args ...interface{}) interface{}
}
