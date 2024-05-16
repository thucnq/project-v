// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: write_ticket_node_participant.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createTicketNodeParticipant = `-- name: CreateTicketNodeParticipant :execresult
insert ignore into ticket_node_participants (workspace_id, ticket_id, node_id,
                                      participant_id, role_type,
                                      participant_type, created_at, updated_at)
values (?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateTicketNodeParticipantParams struct {
	WorkspaceID     string
	TicketID        int64
	NodeID          int64
	ParticipantID   string
	RoleType        int8
	ParticipantType int8
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (q *Queries) CreateTicketNodeParticipant(ctx context.Context, arg CreateTicketNodeParticipantParams) (sql.Result, error) {
	return q.exec(ctx, q.createTicketNodeParticipantStmt, createTicketNodeParticipant,
		arg.WorkspaceID,
		arg.TicketID,
		arg.NodeID,
		arg.ParticipantID,
		arg.RoleType,
		arg.ParticipantType,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const removeTicketNodeParticipant = `-- name: RemoveTicketNodeParticipant :execresult
delete from ticket_node_participants
where workspace_id = ? and ticket_id = ? and node_id = ? and participant_id = ? and role_type = ? and participant_type = ?
`

type RemoveTicketNodeParticipantParams struct {
	WorkspaceID     string
	TicketID        int64
	NodeID          int64
	ParticipantID   string
	RoleType        int8
	ParticipantType int8
}

func (q *Queries) RemoveTicketNodeParticipant(ctx context.Context, arg RemoveTicketNodeParticipantParams) (sql.Result, error) {
	return q.exec(ctx, q.removeTicketNodeParticipantStmt, removeTicketNodeParticipant,
		arg.WorkspaceID,
		arg.TicketID,
		arg.NodeID,
		arg.ParticipantID,
		arg.RoleType,
		arg.ParticipantType,
	)
}
