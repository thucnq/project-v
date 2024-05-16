package serrors

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"project-v/internal/pkg/ccontext"
	"project-v/pkg/container"
	"project-v/pkg/errors"
)

var mapFieldCode = map[string]errors.Code{}

func WrapI18n(
	ctx context.Context, code errors.Code, message string, errs ...error,
) errors.IError {
	lang := ccontext.GetLang(ctx)
	var bundle, _ = container.Resolver[*i18n.Bundle]()
	if bundle == nil {
		return errors.ErrorTraceCtx(ctx, code, message, errs...)
	}

	localizer := i18n.NewLocalizer(bundle, lang)

	msg, err := localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    message,
				Other: message,
			},
		},
	)
	if err != nil {
		msg = message
	}
	return errors.ErrorTraceCtx(ctx, code, msg, errs...)
}

func ErrUnauthenticated(ctx context.Context, err error) errors.IError {
	return WrapI18n(ctx, errors.Unauthenticated, "ErrUnauthenticated", err)
}

func ErrInternal(ctx context.Context, err error) errors.IError {
	return WrapI18n(ctx, errors.Internal, "ErrInternal", err)
}

func ErrFailedPrecondition(ctx context.Context, err error) errors.IError {
	return WrapI18n(ctx, errors.FailedPrecondition, "FailedPrecondition", err)
}

func ErrInvalidArgument(ctx context.Context, err error) errors.IError {
	return WrapI18n(ctx, errors.InvalidArgument, "ErrInvalidArgument", err)
}

func ErrPermissionDenied(ctx context.Context, err error) errors.IError {
	return WrapI18n(ctx, errors.PermissionDenied, "ErrPermissionDenied", err)
}

func ErrTooManyRequests(ctx context.Context, err error) errors.IError {
	return WrapI18n(ctx, errors.ResourceExhausted, "ErrTooManyRequests", err)
}

func ErrInvalidArgumentFromValidate(
	ctx context.Context, err error,
) errors.IError {
	code := errors.InvalidArgument
	msg := "ErrInvalidArgument"

	// try get field and message from err
	var errField string
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		errField = errs[0].Field()
		ns := errs[0].Namespace()
		arr := strings.Split(ns, ".")
		if len(arr) == 3 {
			arr = arr[1:]
		}
		isArr := strings.Index(arr[0], "[")
		if isArr > -1 {
			errField = arr[0][:isArr] + "." + arr[1]
		}
	}

	xerr, ok := err.(*errors.ValidateError)
	if ok {
		msg = xerr.Err
	}

	if len(errField) != 0 {
		xcode, found := mapFieldCode[errField]
		if found {
			code = xcode
		}
	}

	return WrapI18n(ctx, code, msg, fmt.Errorf(err.Error()))
}

func ErrNotFound(ctx context.Context, err error) errors.IError {
	return WrapI18n(ctx, errors.NotFound, "ErrNotFound", err)
}
func HandleError(ctx context.Context, err error) errors.IError {
	if xerr, ok := err.(*errors.APIError); ok {
		msg := xerr.Message
		if len(xerr.Original) > 0 {
			msg = xerr.Original
		}
		return WrapI18n(ctx, xerr.Code, msg, err)
	}
	return WrapI18n(ctx, errors.Internal, "Internal error", err)
}
