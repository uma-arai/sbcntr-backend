package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/horsewin/echo-playground-v2/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// DB ...
type DB struct {
	Host     string
	Username string
	Password string
	DBName   string
	Connect  *sqlx.DB
}

// SQLHandler ... SQL handler struct
type SQLHandler struct {
	Conn *sqlx.DB
}

var (
	sqlHandlerInstance *SQLHandler
	once               sync.Once
)

const (
	dbType = "postgres"
)

// NewSQLHandler ...
func NewSQLHandler() *SQLHandler {
	once.Do(func() {
		c := utils.NewConfigDB()
		USER := c.Postgres.Username
		PASS := c.Postgres.Password
		DBNAME := c.Postgres.DBName
		dbPort := os.Getenv("DB_PORT")
		if dbPort == "" {
			dbPort = "5432"
		}
		PROTOCOL := "host=" + os.Getenv("DB_HOST") + " port=" + dbPort

		// localhostのDBの場合はSSLを無効化
		var sslModeValue string
		if os.Getenv("DB_HOST") == "localhost" {
			sslModeValue = "disable"
		} else {
			sslModeValue = "require" // 本番環境ではSSLを有効にする
		}

		CONNECT := "user=" + USER + " password=" + PASS + " " + PROTOCOL + " dbname=" + DBNAME + " sslmode=" + sslModeValue

		// 標準のSQLコネクションを作成
		db, err := sql.Open(dbType, CONNECT)
		if err != nil {
			log.Fatalf("Error: No database connection established: %v", err)
		}
		conn := sqlx.NewDb(db, dbType)
		err = conn.Ping()
		if err != nil {
			db.Close()
			log.Fatalf("Error: No database connection established: %v", err)
		}

		// 接続成功
		log.Println("DB connected successfully")

		sqlHandlerInstance = &SQLHandler{Conn: conn}
	})

	return sqlHandlerInstance
}

// Where ...
func (handler *SQLHandler) Where(ctx context.Context, out interface{}, table string, whereClause string, whereArgs map[string]interface{}) error {
	// スパンを作成
	tracer := otel.Tracer("sql-handler")
	ctx, span := tracer.Start(ctx, "SQLHandler.Where",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer span.End()

	query := fmt.Sprintf("SELECT * FROM %s", table)
	if whereClause != "" {
		query += fmt.Sprintf(" WHERE %s", whereClause)
	}

	// 属性を追加
	span.SetAttributes(
		attribute.String("db.system", dbType),
		attribute.String("db.statement", query),
		attribute.String("db.operation", "SELECT"),
		attribute.String("db.sql.table", table),
	)

	stmt, err := handler.Conn.PrepareNamed(query)
	if err != nil {
		span.RecordError(err)
		return err
	}

	err = stmt.SelectContext(ctx, out, whereArgs)
	if err != nil {
		span.RecordError(err)
	}
	return err
}

// Scan ...
func (handler *SQLHandler) Scan(ctx context.Context, out interface{}, table string, order string) error {
	// スパンを作成
	tracer := otel.Tracer("sql-handler")
	ctx, span := tracer.Start(ctx, "SQLHandler.Scan",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer span.End()

	query := fmt.Sprintf("SELECT * FROM %s ORDER BY %s;", table, order)

	// 属性を追加
	span.SetAttributes(
		attribute.String("db.system", dbType),
		attribute.String("db.statement", query),
		attribute.String("db.operation", "SELECT"),
		attribute.String("db.sql.table", table),
	)

	err := handler.Conn.SelectContext(ctx, out, query)
	if err != nil {
		span.RecordError(err)
	}
	return err
}

// Count ...
func (handler *SQLHandler) Count(ctx context.Context, out *int, table string, whereClause string, whereArgs map[string]interface{}) error {
	// スパンを作成
	tracer := otel.Tracer("sql-handler")
	ctx, span := tracer.Start(ctx, "SQLHandler.Count",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer span.End()

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)
	if whereClause != "" {
		query += fmt.Sprintf(" WHERE %s", whereClause)
	}

	// 属性を追加
	span.SetAttributes(
		attribute.String("db.system", dbType),
		attribute.String("db.statement", query),
		attribute.String("db.operation", "SELECT"),
		attribute.String("db.sql.table", table),
	)

	var count int
	stmt, err := handler.Conn.PrepareNamed(query)
	if err != nil {
		span.RecordError(err)
		return err
	}

	err = stmt.GetContext(ctx, &count, whereArgs)
	*out = count
	if err != nil {
		span.RecordError(err)
	}
	return err
}

// Create ...
func (handler *SQLHandler) Create(ctx context.Context, input map[string]interface{}, table string) error {
	// スパンを作成
	tracer := otel.Tracer("sql-handler")
	ctx, span := tracer.Start(ctx, "SQLHandler.Create",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer span.End()

	// カラム名とプレースホルダーを構築
	columns := make([]string, 0)
	placeholders := make([]string, 0)

	// inputのキーと値をそれぞれ列と値に追加
	for key := range input {
		// IDが設定されていても無視する
		if key == "id" {
			continue
		}

		columns = append(columns, key)
		placeholders = append(placeholders, fmt.Sprintf(":%s", key)) // プレースホルダーを使う
	}

	// クエリを構築
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, strings.Join(columns, ","), strings.Join(placeholders, ","))

	// 属性を追加
	span.SetAttributes(
		attribute.String("db.system", dbType),
		attribute.String("db.statement", query),
		attribute.String("db.operation", "INSERT"),
		attribute.String("db.sql.table", table),
	)

	_, err := handler.Conn.NamedExecContext(ctx, query, input)
	if err != nil {
		span.RecordError(err)
	}
	return err
}

// Update ...
func (handler *SQLHandler) Update(ctx context.Context, setParams map[string]interface{}, table string, whereClause string, whereParams map[string]interface{}) error {
	// スパンを作成
	tracer := otel.Tracer("sql-handler")
	ctx, span := tracer.Start(ctx, "SQLHandler.Update",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer span.End()

	// SET句を構築（SET用パラメータのみ使用）
	setColumns, setPlaceholders, _ := buildNamedParameters(setParams)
	setClauses := make([]string, len(setColumns))
	for i, col := range setColumns {
		setClauses[i] = fmt.Sprintf("%s = %s", col, setPlaceholders[i])
	}

	// WHERE句のパラメータ名の重複を回避
	adjustedWhereClause := whereClause
	adjustedWhereParams := make(map[string]interface{})

	for key, value := range whereParams {
		// SET用パラメータと重複する場合は接尾辞を付ける
		if _, exists := setParams[key]; exists {
			newKey := key + "_where"
			adjustedWhereParams[newKey] = value
			// WHERE句内のプレースホルダーも置換
			adjustedWhereClause = strings.ReplaceAll(adjustedWhereClause, ":"+key, ":"+newKey)
		} else {
			adjustedWhereParams[key] = value
		}
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, strings.Join(setClauses, ","), adjustedWhereClause)

	// 実行用にパラメータをマージ（重複回避済み）
	allParams := make(map[string]interface{})
	for k, v := range setParams {
		allParams[k] = v
	}
	for k, v := range adjustedWhereParams {
		allParams[k] = v
	}

	// 属性を追加
	span.SetAttributes(
		attribute.String("db.system", dbType),
		attribute.String("db.statement", query),
		attribute.String("db.operation", "UPDATE"),
		attribute.String("db.sql.table", table),
	)

	_, err := handler.Conn.NamedExecContext(ctx, query, allParams)

	if err != nil {
		span.RecordError(err)
	}

	return err
}

// Delete ...
func (handler *SQLHandler) Delete(ctx context.Context, in map[string]interface{}, table string) error {
	// スパンを作成
	tracer := otel.Tracer("sql-handler")
	ctx, span := tracer.Start(ctx, "SQLHandler.Delete",
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer span.End()

	columns, _, values := buildNamedParameters(in)

	whereClauses := make([]string, len(columns))
	for i, col := range columns {
		whereClauses[i] = fmt.Sprintf("%s = %v", col, values[col])
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE %s", table, strings.Join(whereClauses, ","))

	// 属性を追加
	span.SetAttributes(
		attribute.String("db.system", dbType),
		attribute.String("db.statement", query),
		attribute.String("db.operation", "DELETE"),
		attribute.String("db.sql.table", table),
	)

	_, err := handler.Conn.NamedExecContext(ctx, query, values)
	if err != nil {
		span.RecordError(err)
	}

	return err
}

func buildNamedParameters(input map[string]interface{}) (columns []string, placeholderNames []string, values map[string]interface{}) {
	columns = []string{}
	placeholderNames = []string{}
	values = make(map[string]interface{})

	for key, value := range input {
		if value != nil {
			columns = append(columns, key)
			placeholderNames = append(placeholderNames, ":"+key)
			values[key] = value
		}
	}

	return
}
