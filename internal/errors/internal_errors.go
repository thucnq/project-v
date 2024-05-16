package serrors

import (
	"project-v/pkg/errors"
)

var (
	InvalidObjID      = errors.Error(errors.InvalidArgument, "invalid id")
	ErrParentNotFound = errors.Error(
		errors.InvalidArgument, "ErrParentNotFound",
	)
	ErrDuplicateEntry = errors.Error(
		errors.InvalidArgument, "ErrDuplicateEntry",
	)
	ResourceNotFound = errors.Error(
		errors.NotFound, "ResourceNotFound",
	)
)
