// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: write_ticket.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const updateStatusTicket = `-- name: UpdateStatusTicket :execresult
UPDATE tickets
SET
    status = ?,
    updated_by = ?,
    updated_at = ?
WHERE workspace_id = ?
  AND id = ?
`

type UpdateStatusTicketParams struct {
	Status      int8
	UpdatedBy   int64
	UpdatedAt   time.Time
	WorkspaceID string
	ID          int64
}

func (q *Queries) UpdateStatusTicket(ctx context.Context, arg UpdateStatusTicketParams) (sql.Result, error) {
	return q.exec(ctx, q.updateStatusTicketStmt, updateStatusTicket,
		arg.Status,
		arg.UpdatedBy,
		arg.UpdatedAt,
		arg.WorkspaceID,
		arg.ID,
	)
}
