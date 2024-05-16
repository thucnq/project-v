// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: delete_ticket_notification.sql

package sqlc

import (
	"context"
	"database/sql"
)

const deleteTicketNotification = `-- name: DeleteTicketNotification :execresult
DELETE FROM ticket_notifications WHERE workspace_id=? AND id=?
`

type DeleteTicketNotificationParams struct {
	WorkspaceID string
	ID          int64
}

func (q *Queries) DeleteTicketNotification(ctx context.Context, arg DeleteTicketNotificationParams) (sql.Result, error) {
	return q.exec(ctx, q.deleteTicketNotificationStmt, deleteTicketNotification, arg.WorkspaceID, arg.ID)
}
