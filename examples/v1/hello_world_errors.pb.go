// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	context "context"
	fmt "fmt"
	errors "github.com/102er/protoc-gen-gogo-errors/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// is i18n value

var (
	zhMap map[string]string
	enMap map[string]string
)

func init() {
	zhMap = make(map[string]string)

	zhMap[ErrorReason_USER_NOT_FOUND.String()] = "404zh"
	zhMap[ErrorReason_USER_NOT_FOUND_t.String()] = "40401zh"
}
func init() {
	enMap = make(map[string]string)

	enMap[ErrorReason_USER_NOT_FOUND.String()] = "404en"
	enMap[ErrorReason_USER_NOT_FOUND_t.String()] = "40401en"
}
func getMessage(ctx context.Context, key string) string {
	l := "en"
	if ll, ok := errors.FromErrorLangCtx(ctx); ok && ll == "zh" {
		l = "zh"
	}
	if l == "zh" {
		return zhMap[key]
	}
	return enMap[key]
}

func IsUserNotFound(ctx context.Context, err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)

	return e.Message == getMessage(ctx, ErrorReason_USER_NOT_FOUND.String()) && e.Code == 404
}
func ErrorUserNotFound(ctx context.Context, format string, args ...interface{}) *errors.Error {
	return errors.NewWithErrCode(404, 1231, getMessage(ctx, ErrorReason_USER_NOT_FOUND.String()), fmt.Sprintf(format, args...))
}
func IsUserNotFoundT(ctx context.Context, err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)

	return e.Message == getMessage(ctx, ErrorReason_USER_NOT_FOUND_t.String()) && e.Code == 403
}
func ErrorUserNotFoundT(ctx context.Context, format string, args ...interface{}) *errors.Error {
	return errors.NewWithErrCode(403, 41204, getMessage(ctx, ErrorReason_USER_NOT_FOUND_t.String()), fmt.Sprintf(format, args...))
}