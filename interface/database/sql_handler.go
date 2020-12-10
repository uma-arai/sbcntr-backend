package database

// SQLHandler ...
type SQLHandler interface {
	Where(out interface{}, query interface{}, args ...interface{}) interface{}
}
