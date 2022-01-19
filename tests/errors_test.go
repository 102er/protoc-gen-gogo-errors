package tests

import (
	"context"
	"github.com/102er/protoc-gen-gogo-errors/errors"
	v1 "github.com/102er/protoc-gen-gogo-errors/examples/v1"
	"testing"
)
import _ "github.com/102er/protoc-gen-gogo-errors/examples/v1"

func TestUserNotFound(t *testing.T) {
	ctx := context.Background()
	errors.NewErrorLangCtx(ctx, "en")
	e := v1.ErrorUserNotFound(ctx, "test")
	t.Log(e)
}
