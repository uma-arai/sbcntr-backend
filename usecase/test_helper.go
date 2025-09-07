package usecase

import (
	"context"
)

func init() {
}

// testContext はテスト用のコンテキストを作成する
func testContext() context.Context {
	return context.Background()
}
