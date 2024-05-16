// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: delete_advance_setting_participants.sql

package sqlc

import (
	"context"
	"database/sql"
)

const deleteAdvanceSettingParticipants = `-- name: DeleteAdvanceSettingParticipants :execresult
DELETE FROM advance_setting_participants WHERE workspace_id=? AND ref_id=?
`

type DeleteAdvanceSettingParticipantsParams struct {
	WorkspaceID string
	RefID       int64
}

func (q *Queries) DeleteAdvanceSettingParticipants(ctx context.Context, arg DeleteAdvanceSettingParticipantsParams) (sql.Result, error) {
	return q.exec(ctx, q.deleteAdvanceSettingParticipantsStmt, deleteAdvanceSettingParticipants, arg.WorkspaceID, arg.RefID)
}
