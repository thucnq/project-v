// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: insert_advance_setting.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createAdvanceSetting = `-- name: CreateAdvanceSetting :execresult
INSERT INTO advance_settings (
	workspace_id,
	workflow_id,
	is_published,
	created_by,
	updated_by,
	admin_option_participant_value,
	admin_option_participant_ref,
	allow_cancel_value,
	allow_cancel_participant_value,
	allow_cancel_participant_ref,
	allow_delete_value,
	allow_delete_participant_value,
	allow_delete_participant_ref,
	close_option_value,
	close_option_participant_value,
	close_option_participant_ref,
	approve_before_on_hold_value,
	approve_before_on_hold_participant_value,
	approve_before_on_hold_participant_ref,
	rate_option_value,
	version
)
VALUES (
	?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`

type CreateAdvanceSettingParams struct {
	WorkspaceID                         string
	WorkflowID                          int64
	IsPublished                         bool
	CreatedBy                           int64
	UpdatedBy                           int64
	AdminOptionParticipantValue         int32
	AdminOptionParticipantRef           int64
	AllowCancelValue                    int32
	AllowCancelParticipantValue         int32
	AllowCancelParticipantRef           int64
	AllowDeleteValue                    int32
	AllowDeleteParticipantValue         int32
	AllowDeleteParticipantRef           int64
	CloseOptionValue                    int32
	CloseOptionParticipantValue         int32
	CloseOptionParticipantRef           int64
	ApproveBeforeOnHoldValue            int32
	ApproveBeforeOnHoldParticipantValue int32
	ApproveBeforeOnHoldParticipantRef   int64
	RateOptionValue                     int32
	Version                             int32
}

func (q *Queries) CreateAdvanceSetting(ctx context.Context, arg CreateAdvanceSettingParams) (sql.Result, error) {
	return q.exec(ctx, q.createAdvanceSettingStmt, createAdvanceSetting,
		arg.WorkspaceID,
		arg.WorkflowID,
		arg.IsPublished,
		arg.CreatedBy,
		arg.UpdatedBy,
		arg.AdminOptionParticipantValue,
		arg.AdminOptionParticipantRef,
		arg.AllowCancelValue,
		arg.AllowCancelParticipantValue,
		arg.AllowCancelParticipantRef,
		arg.AllowDeleteValue,
		arg.AllowDeleteParticipantValue,
		arg.AllowDeleteParticipantRef,
		arg.CloseOptionValue,
		arg.CloseOptionParticipantValue,
		arg.CloseOptionParticipantRef,
		arg.ApproveBeforeOnHoldValue,
		arg.ApproveBeforeOnHoldParticipantValue,
		arg.ApproveBeforeOnHoldParticipantRef,
		arg.RateOptionValue,
		arg.Version,
	)
}
