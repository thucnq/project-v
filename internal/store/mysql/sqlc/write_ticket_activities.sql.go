// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: write_ticket_activities.sql

package sqlc

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

const createTicketActivity = `-- name: CreateTicketActivity :execresult
INSERT INTO ticket_activities (
    workspace_id, ticket_id, id, actor_id, actor_type, type,
   target_id, target_type, object_id, object_type, created_at, content
) VALUES (
        ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`

type CreateTicketActivityParams struct {
	WorkspaceID string
	TicketID    int64
	ID          int64
	ActorID     string
	ActorType   int8
	Type        int8
	TargetID    string
	TargetType  int8
	ObjectID    string
	ObjectType  int8
	CreatedAt   time.Time
	Content     json.RawMessage
}

func (q *Queries) CreateTicketActivity(ctx context.Context, arg CreateTicketActivityParams) (sql.Result, error) {
	return q.exec(ctx, q.createTicketActivityStmt, createTicketActivity,
		arg.WorkspaceID,
		arg.TicketID,
		arg.ID,
		arg.ActorID,
		arg.ActorType,
		arg.Type,
		arg.TargetID,
		arg.TargetType,
		arg.ObjectID,
		arg.ObjectType,
		arg.CreatedAt,
		arg.Content,
	)
}
