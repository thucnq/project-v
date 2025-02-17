// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: select_advance_setting.sql

package sqlc

import (
	"context"
)

const getAdvanceSetting = `-- name: GetAdvanceSetting :one
SELECT created_by, updated_by, created_at, updated_at, workspace_id, workflow_id, is_published, version, admin_option_participant_value, admin_option_participant_ref, allow_cancel_value, allow_cancel_participant_value, allow_cancel_participant_ref, allow_delete_value, allow_delete_participant_value, allow_delete_participant_ref, close_option_value, close_option_participant_value, close_option_participant_ref, rate_option_value, approve_before_on_hold_value, approve_before_on_hold_participant_value, approve_before_on_hold_participant_ref
FROM advance_settings
WHERE workspace_id=? AND workflow_id=? LIMIT 1
`

type GetAdvanceSettingParams struct {
	WorkspaceID string
	WorkflowID  int64
}

func (q *Queries) GetAdvanceSetting(ctx context.Context, arg GetAdvanceSettingParams) (AdvanceSetting, error) {
	row := q.queryRow(ctx, q.getAdvanceSettingStmt, getAdvanceSetting, arg.WorkspaceID, arg.WorkflowID)
	var i AdvanceSetting
	err := row.Scan(
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.WorkspaceID,
		&i.WorkflowID,
		&i.IsPublished,
		&i.Version,
		&i.AdminOptionParticipantValue,
		&i.AdminOptionParticipantRef,
		&i.AllowCancelValue,
		&i.AllowCancelParticipantValue,
		&i.AllowCancelParticipantRef,
		&i.AllowDeleteValue,
		&i.AllowDeleteParticipantValue,
		&i.AllowDeleteParticipantRef,
		&i.CloseOptionValue,
		&i.CloseOptionParticipantValue,
		&i.CloseOptionParticipantRef,
		&i.RateOptionValue,
		&i.ApproveBeforeOnHoldValue,
		&i.ApproveBeforeOnHoldParticipantValue,
		&i.ApproveBeforeOnHoldParticipantRef,
	)
	return i, err
}
