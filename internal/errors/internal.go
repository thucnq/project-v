package serrors

import (
	"errors"
)

var (
	ErrDataNotFound        = errors.New("not found")
	ErrDataDuplicate       = errors.New("duplicate")
	ErrInvalidWorkspaceID  = errors.New("invalid workspace id")
	ErrInvalidDepartmentID = errors.New("invalid department id")
	ErrInvalidRoleID       = errors.New("invalid role id")
	ErrWorkspaceIsSyncing  = errors.New("workspace is syncing")
)
