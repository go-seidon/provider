package system

import (
	"github.com/go-seidon/provider/status"
	"google.golang.org/grpc/codes"
	statuses "google.golang.org/grpc/status"
)

func FromGrpc(err error) *Error {
	code := status.ACTION_FAILED
	message := err.Error()

	gstatus, ok := statuses.FromError(err)
	if !ok {
		return &Error{
			Code:    code,
			Message: message,
		}
	}

	message = gstatus.Message()
	switch gstatus.Code() {
	case codes.Internal:
		code = status.ACTION_FAILED //explicit setup
		message = "failed processing request"
	case codes.InvalidArgument:
		code = status.INVALID_PARAM
	case codes.PermissionDenied:
		code = status.ACTION_FORBIDDEN
	case codes.NotFound:
		code = status.RESOURCE_NOTFOUND
	case codes.Unavailable:
		code = status.COMMUNICATION_FAILURE
		message = "failed communicating with the host party"
	case codes.AlreadyExists:
		code = status.RESOURCE_EXISTS
	case codes.Unauthenticated:
		code = status.INVALID_CREDENTIAL
	}
	return &Error{
		Code:    code,
		Message: message,
	}
}
