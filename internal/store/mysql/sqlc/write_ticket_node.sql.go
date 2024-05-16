// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: write_ticket_node.sql

package sqlc

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

const createTicketNode = `-- name: CreateTicketNode :execresult
insert into ticket_nodes (workspace_id, assignee_id, ticket_id, node_id, status,
                          type, sla, ` + "`" + `option` + "`" + `, deadline, created_at, updated_at)
values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateTicketNodeParams struct {
	WorkspaceID string
	AssigneeID  int64
	TicketID    int64
	NodeID      int64
	Status      int8
	Type        int8
	Sla         json.RawMessage
	Option      json.RawMessage
	Deadline    sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreateTicketNode(ctx context.Context, arg CreateTicketNodeParams) (sql.Result, error) {
	return q.exec(ctx, q.createTicketNodeStmt, createTicketNode,
		arg.WorkspaceID,
		arg.AssigneeID,
		arg.TicketID,
		arg.NodeID,
		arg.Status,
		arg.Type,
		arg.Sla,
		arg.Option,
		arg.Deadline,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const resetStatusAndDeadlineTicketNodesByTicketID = `-- name: ResetStatusAndDeadlineTicketNodesByTicketID :execresult
UPDATE ticket_nodes
SET deadline = NULL,
    status = 1,
    updated_at = ?
WHERE workspace_id = ?
  AND ticket_id = ?
  AND status <> 1
`

type ResetStatusAndDeadlineTicketNodesByTicketIDParams struct {
	UpdatedAt   time.Time
	WorkspaceID string
	TicketID    int64
}

func (q *Queries) ResetStatusAndDeadlineTicketNodesByTicketID(ctx context.Context, arg ResetStatusAndDeadlineTicketNodesByTicketIDParams) (sql.Result, error) {
	return q.exec(ctx, q.resetStatusAndDeadlineTicketNodesByTicketIDStmt, resetStatusAndDeadlineTicketNodesByTicketID, arg.UpdatedAt, arg.WorkspaceID, arg.TicketID)
}

const responseTicketNode = `-- name: ResponseTicketNode :execresult
update ticket_nodes
set first_response_at = ?, updated_at = ?
where workspace_id = ? and ticket_id = ? and node_id = ? and first_response_at is null
`

type ResponseTicketNodeParams struct {
	FirstResponseAt sql.NullTime
	UpdatedAt       time.Time
	WorkspaceID     string
	TicketID        int64
	NodeID          int64
}

func (q *Queries) ResponseTicketNode(ctx context.Context, arg ResponseTicketNodeParams) (sql.Result, error) {
	return q.exec(ctx, q.responseTicketNodeStmt, responseTicketNode,
		arg.FirstResponseAt,
		arg.UpdatedAt,
		arg.WorkspaceID,
		arg.TicketID,
		arg.NodeID,
	)
}

const updateStatusTicketNode = `-- name: UpdateStatusTicketNode :execresult
UPDATE ticket_nodes
SET
    status = ?,
    updated_at = ?
WHERE workspace_id = ?
  AND ticket_id = ?
  AND node_id = ?
`

type UpdateStatusTicketNodeParams struct {
	Status      int8
	UpdatedAt   time.Time
	WorkspaceID string
	TicketID    int64
	NodeID      int64
}

func (q *Queries) UpdateStatusTicketNode(ctx context.Context, arg UpdateStatusTicketNodeParams) (sql.Result, error) {
	return q.exec(ctx, q.updateStatusTicketNodeStmt, updateStatusTicketNode,
		arg.Status,
		arg.UpdatedAt,
		arg.WorkspaceID,
		arg.TicketID,
		arg.NodeID,
	)
}

const updateTicketNodeCancelOnHold = `-- name: UpdateTicketNodeCancelOnHold :execresult
UPDATE ticket_nodes
SET
    status = ?,
    deadline = ?,
    updated_at = ?
WHERE workspace_id = ?
  AND ticket_id = ?
  AND node_id = ?
`

type UpdateTicketNodeCancelOnHoldParams struct {
	Status      int8
	Deadline    sql.NullTime
	UpdatedAt   time.Time
	WorkspaceID string
	TicketID    int64
	NodeID      int64
}

func (q *Queries) UpdateTicketNodeCancelOnHold(ctx context.Context, arg UpdateTicketNodeCancelOnHoldParams) (sql.Result, error) {
	return q.exec(ctx, q.updateTicketNodeCancelOnHoldStmt, updateTicketNodeCancelOnHold,
		arg.Status,
		arg.Deadline,
		arg.UpdatedAt,
		arg.WorkspaceID,
		arg.TicketID,
		arg.NodeID,
	)
}

const updateTicketNodeToOnHold = `-- name: UpdateTicketNodeToOnHold :execresult
UPDATE ticket_nodes
SET deadline = NULL,
    status = ?,
    updated_at = ?
WHERE workspace_id = ?
  AND ticket_id = ?
  AND node_id = ?
`

type UpdateTicketNodeToOnHoldParams struct {
	Status      int8
	UpdatedAt   time.Time
	WorkspaceID string
	TicketID    int64
	NodeID      int64
}

func (q *Queries) UpdateTicketNodeToOnHold(ctx context.Context, arg UpdateTicketNodeToOnHoldParams) (sql.Result, error) {
	return q.exec(ctx, q.updateTicketNodeToOnHoldStmt, updateTicketNodeToOnHold,
		arg.Status,
		arg.UpdatedAt,
		arg.WorkspaceID,
		arg.TicketID,
		arg.NodeID,
	)
}
