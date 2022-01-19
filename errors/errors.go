package errors

import (
	"context"
	"errors"
	"fmt"
	httpstatus "github.com/go-kratos/kratos/v2/transport/http/status"
	"github.com/gogo/protobuf/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

const (
	// UnknownCode is unknown code for error info.
	UnknownCode = 500
	// UnknownReason is unknown reason for error info.
	UnknownReason = ""
	// SupportPackageIsVersion1 this constant should not be referenced by any other code.
	SupportPackageIsVersion1 = true
)

//go:generate protoc -I. --go_out=paths=source_relative:. errors.proto

func (x *Error) Error() string {
	return fmt.Sprintf("error: code = %d  message = %s detail = %s metadata = %v", x.Code, x.Message, x.Detail, x.Metadata)
}

// GRPCStatus returns the Status represented by se.
func (x *Error) GRPCStatus() *status.Status {
	s, _ := status.New(httpstatus.ToGRPCCode(int(x.Code)), x.Message).
		WithDetails(&errdetails.ErrorInfo{
			Reason:   x.Message,
			Metadata: x.Metadata,
		})
	return s
}

// Is matches each error in the chain with the target value.
func (x *Error) Is(err error) bool {
	if se := new(Error); errors.As(err, &se) {
		return se.Message == x.Message
	}
	return false
}

// WithMetadata with an MD formed by the mapping of key, value.
func (x *Error) WithMetadata(md map[string]string) *Error {
	err := proto.Clone(x).(*Error)
	err.Metadata = md
	return err
}

// New returns an error object for the code, message.
func New(code int, message, detail string) *Error {
	return &Error{
		Code:    int32(code),
		Message: message,
		Detail:  detail,
	}
}

// Newf New(code fmt.Sprintf(format, a...))
func Newf(code int, reason, format string, a ...interface{}) *Error {
	return New(code, reason, fmt.Sprintf(format, a...))
}

// Errorf returns an error object for the code, message and error info.
func Errorf(code int, reason, format string, a ...interface{}) error {
	return New(code, reason, fmt.Sprintf(format, a...))
}

// NewWithErrCode returns an error object for the code, message.
func NewWithErrCode(code, eCode int, message, detail string) *Error {
	return &Error{
		Code:    int32(code),
		ErrCode: int32(eCode),
		Message: message,
		Detail:  detail,
	}
}

// Code returns the http code for a error.
// It supports wrapped errors.
func Code(err error) int {
	if err == nil {
		return 200 //nolint:gomnd
	}
	return int(FromError(err).Code)
}

// Message returns the reason for a particular error.
// It supports wrapped errors.
func Message(err error) string {
	if err == nil {
		return UnknownReason
	}
	return FromError(err).Message
}

// FromError try to convert an error to *Error.
// It supports wrapped errors.
func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if se := new(Error); errors.As(err, &se) {
		return se
	}
	gs, ok := status.FromError(err)
	if ok {
		ret := New(
			httpstatus.FromGRPCCode(gs.Code()),
			UnknownReason,
			gs.Message(),
		)
		for _, detail := range gs.Details() {
			switch d := detail.(type) {
			case *errdetails.ErrorInfo:
				ret.Message = d.Reason
				return ret.WithMetadata(d.Metadata)
			}
		}
		return ret
	}
	return New(UnknownCode, UnknownReason, err.Error())
}

type ErrorLangCtx struct{}

func NewErrorLangCtx(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, ErrorLangCtx{}, lang)
}
func FromErrorLangCtx(ctx context.Context) (string, bool) {
	l, ok := ctx.Value(ErrorLangCtx{}).(string)
	return l, ok
}
