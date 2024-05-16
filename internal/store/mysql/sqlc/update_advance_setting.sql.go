// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: update_advance_setting.sql

package sqlc

import (
	"context"
	"database/sql"
)

const updateAdvanceSetting = `-- name: UpdateAdvanceSetting :execresult
UPDATE advance_settings SET
	is_published=?,
	allow_cancel_value=?,
	allow_cancel_participant_value=?,
	allow_delete_value=?,
	allow_delete_participant_value=?,
	close_option_value=?,
	close_option_participant_value=?,
	approve_before_on_hold_value=?,
	approve_before_on_hold_participant_value=?,
	rate_option_value=?,
	version=?,
	updated_by=?,
	updated_at=?
WHERE workspace_id=? AND workflow_id=?
`

type UpdateAdvanceSettingParams struct {
	IsPublished                         bool
	AllowCancelValue                    int32
	AllowCancelParticipantValue         int32
	AllowDeleteValue                    int32
	AllowDeleteParticipantValue         int32
	CloseOptionValue                    int32
	CloseOptionParticipantValue         int32
	ApproveBeforeOnHoldValue            int32
	ApproveBeforeOnHoldParticipantValue int32
	RateOptionValue                     int32
	Version                             int32
	UpdatedBy                           int64
	UpdatedAt                           sql.NullTime
	WorkspaceID                         string
	WorkflowID                          int64
}

func (q *Queries) UpdateAdvanceSetting(ctx context.Context, arg UpdateAdvanceSettingParams) (sql.Result, error) {
	return q.exec(ctx, q.updateAdvanceSettingStmt, updateAdvanceSetting,
		arg.IsPublished,
		arg.AllowCancelValue,
		arg.AllowCancelParticipantValue,
		arg.AllowDeleteValue,
		arg.AllowDeleteParticipantValue,
		arg.CloseOptionValue,
		arg.CloseOptionParticipantValue,
		arg.ApproveBeforeOnHoldValue,
		arg.ApproveBeforeOnHoldParticipantValue,
		arg.RateOptionValue,
		arg.Version,
		arg.UpdatedBy,
		arg.UpdatedAt,
		arg.WorkspaceID,
		arg.WorkflowID,
	)
}
