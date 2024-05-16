package ccontext

import (
	"context"
	"errors"
	"strconv"
	"time"

	"project-v/internal/constant"
)

const (
	XGapoRole        = "x-gapo-role"
	XGapoUserId      = "x-gapo-user-id"
	XGapoApiKey      = "x-gapo-api-key"
	XGapoLang        = "x-gapo-lang"
	XGapoWorkspaceId = "x-gapo-workspace-id"
)

type loginInfo struct {
	UserID int64

	WorkspaceID string
	Lang        string
	Role        string
}

var (
	ErrNotFound            = errors.New("Not found")
	ErrWorkspaceIDNotFound = errors.New("WorkspaceID not found in header")
	ErrUserIDNotFound      = errors.New("UserID not found in header")
)

// GetWorkspaceID get workspace id from context
func GetWorkspaceID(ctx context.Context) string {
	value := ctx.Value(constant.XGapoWorkspaceId)
	if id, ok := value.(string); ok {
		return id
	}
	return ""
}

// GetUserID get workspace id from context
func GetUserID(ctx context.Context) string {
	value := ctx.Value(constant.XGapoUserId)
	if id, ok := value.(string); ok {
		return id
	}
	return ""
}

// GetRole get workspace id from context
func GetRole(ctx context.Context) string {
	value := ctx.Value(constant.XGapoRole)
	if id, ok := value.(string); ok {
		return id
	}
	return ""
}

// GetApiKey get workspace id from context
func GetApiKey(ctx context.Context) string {
	value := ctx.Value(constant.XGapoApiKey)
	if id, ok := value.(string); ok {
		return id
	}
	return ""
}

func GetLang(ctx context.Context) string {
	value := ctx.Value(constant.XGapoLang)
	id, ok := value.(string)

	if !ok {
		return constant.DefaultLang
	}

	if _, ok := constant.SupportLang[id]; ok {
		return id
	}
	return constant.DefaultLang
}

// GetTimezone get timezone from context
func GetTimezone(ctx context.Context) *time.Location {
	value := ctx.Value(constant.HeaderTimezoneOffset)
	if value == nil {
		return time.FixedZone("UTC+7", 25200)
	}
	tz, err := strconv.Atoi(value.(string))
	if err != nil {
		return time.FixedZone("UTC+7", 25200)
	}
	return time.FixedZone("UTC"+value.(string), tz*60*60)
}

// GetWorkspaceIDV2 get workspace id from context
func GetWorkspaceIDV2(ctx context.Context) (string, error) {
	value := ctx.Value(constant.XGapoWorkspaceId)
	if id, ok := value.(string); ok {
		return id, nil
	}
	return "", ErrWorkspaceIDNotFound
}

// GetUserIDV2 get workspace id from context
func GetUserIDV2(ctx context.Context) (int64, error) {
	rID, err := parseIDToInt64(ctx, XGapoUserId)
	if err != nil {
		err = ErrWorkspaceIDNotFound
	}
	return rID, err
}

func NewLogin(ctx context.Context) (loginInfo, error) {
	var lgInfo loginInfo
	wsID, err := GetWorkspaceIDV2(ctx)
	if err != nil {
		return lgInfo, err
	}

	userID, err := GetUserIDV2(ctx)
	if err != nil {
		return lgInfo, err
	}

	lang := GetLang(ctx)
	role := GetRole(ctx)

	lgInfo.WorkspaceID = wsID
	lgInfo.UserID = userID
	lgInfo.Lang = lang
	lgInfo.Role = role

	return lgInfo, err
}

func parseIDToInt64(ctx context.Context, header string) (int64, error) {
	var (
		rID int64
		err error
	)
	value := ctx.Value(header)
	if id, ok := value.(string); ok {
		rID, err = strconv.ParseInt(id, 10, 64)
	} else {
		err = ErrNotFound
	}
	return rID, err

}
