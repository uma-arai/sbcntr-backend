package database

// SQLHandler ...
type SQLHandler interface {
	Where(out interface{}, query interface{}, args ...interface{}) interface{}
	Scan(out interface{}, order string) interface{}
	Count(out *int, model interface{}) interface{}
}
