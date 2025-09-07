package database

import "context"

// SQLHandler ...
type SQLHandler interface {
	Where(ctx context.Context, out interface{}, tableName string, clause string, args map[string]interface{}) error
	Scan(ctx context.Context, out interface{}, tableName string, order string) error
	Count(ctx context.Context, out *int, tableName string, clause string, args map[string]interface{}) error
	Create(ctx context.Context, in map[string]interface{}, tableName string) error
	Update(ctx context.Context, setParams map[string]interface{}, tableName string, whereClause string, whereParams map[string]interface{}) error
	Delete(ctx context.Context, in map[string]interface{}, tableName string) error
}
