package errors

import (
	"context"
	"errors"
	"net/http"

	"github.com/rs/zerolog"

	"github.com/labstack/echo/v4"
)

// BusinessError インターフェース
type BusinessError interface {
	error
	Code() string
	HTTPStatus() int
	Message(locale string) string
	IsPublic() bool
	OriginalError() error // 元のエラーを取得
}

// businessError 構造体
type businessError struct {
	code        string
	def         messageDef
	originalErr error // 元のエラーを保持
}

func (e *businessError) Error() string {
	if msg, ok := e.def.Message["en"]; ok {
		return msg
	}
	return e.code
}

func (e *businessError) Code() string {
	return e.def.MessageCode
}

func (e *businessError) HTTPStatus() int {
	return e.def.StatusCode
}

func (e *businessError) Message(locale string) string {
	if msg, ok := e.def.Message[locale]; ok {
		return msg
	}
	return e.def.Message["en"]
}

// 5xxは非公開、4xxは公開
func (e *businessError) IsPublic() bool {
	return e.def.StatusCode < 500
}

func (e *businessError) OriginalError() error {
	return e.originalErr
}

// Unwrap errors.Is/As で使用されるUnwrapメソッド
func (e *businessError) Unwrap() error {
	return e.originalErr
}

// NewBusinessError エラー生成関数
func NewBusinessError(code string, originalErr error) BusinessError {
	def, ok := messages[code]
	if !ok {
		// 定義漏れ時のフォールバック
		def = messageDef{
			StatusCode:  500,
			MessageCode: "internal_error",
			Message:     map[string]string{"en": "internal error"},
		}
	}
	return &businessError{code: code, def: def, originalErr: originalErr}
}

// NewEchoHTTPError errorを受け取り、エラーがBusinessErrorかどうかを判定
// 最終的にはecho.NewHTTPErrorを返す
func NewEchoHTTPError(ctx context.Context, err error) *echo.HTTPError {
	var be BusinessError
	if errors.As(err, &be) {
		if be.IsPublic() {
			return echo.NewHTTPError(be.HTTPStatus(), be.Error())
		}
		// 5xxエラーは詳細を隠しつつlogに記録
		if be.OriginalError() != nil {
			zerolog.Ctx(ctx).Error().Err(be).
				Str("error_code", be.Code()).
				Str("origin_error", be.OriginalError().Error()).
				Msg("business logic error")
		}
	}

	// 予期せぬエラー
	return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
}
