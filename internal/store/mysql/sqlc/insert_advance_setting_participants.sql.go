// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: insert_advance_setting_participants.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createAdvanceSettingParticipant = `-- name: CreateAdvanceSettingParticipant :execresult
INSERT IGNORE INTO advance_setting_participants (
	workspace_id,
	ref_id,
	participant_id,
	type
)
VALUES (
	?, ?, ?, ?
)
`

type CreateAdvanceSettingParticipantParams struct {
	WorkspaceID   string
	RefID         int64
	ParticipantID string
	Type          int8
}

func (q *Queries) CreateAdvanceSettingParticipant(ctx context.Context, arg CreateAdvanceSettingParticipantParams) (sql.Result, error) {
	return q.exec(ctx, q.createAdvanceSettingParticipantStmt, createAdvanceSettingParticipant,
		arg.WorkspaceID,
		arg.RefID,
		arg.ParticipantID,
		arg.Type,
	)
}
