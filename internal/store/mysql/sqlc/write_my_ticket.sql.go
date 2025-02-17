// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: write_my_ticket.sql

package sqlc

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

const createMyTicket = `-- name: CreateMyTicket :execresult
INSERT INTO my_tickets (
    workspace_id, workflow_id, id,
    code, title, priority, is_private, status,
    current_node_id, current_node_name, current_node_status, current_node_type,
    node_deadline_at, assignee_id,
    department_id, workflow, tags, ref_ticket_ids,
    rating_point, review,
    latest_recent_activity_user_id, latest_recent_activity_at,
    created_by, updated_by, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateMyTicketParams struct {
	WorkspaceID                string
	WorkflowID                 int64
	ID                         int64
	Code                       string
	Title                      string
	Priority                   int8
	IsPrivate                  bool
	Status                     int8
	CurrentNodeID              int64
	CurrentNodeName            string
	CurrentNodeStatus          int8
	CurrentNodeType            int8
	NodeDeadlineAt             sql.NullTime
	AssigneeID                 sql.NullInt64
	DepartmentID               sql.NullString
	Workflow                   json.RawMessage
	Tags                       json.RawMessage
	RefTicketIds               json.RawMessage
	RatingPoint                int8
	Review                     string
	LatestRecentActivityUserID sql.NullInt64
	LatestRecentActivityAt     sql.NullTime
	CreatedBy                  int64
	UpdatedBy                  int64
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
}

func (q *Queries) CreateMyTicket(ctx context.Context, arg CreateMyTicketParams) (sql.Result, error) {
	return q.exec(ctx, q.createMyTicketStmt, createMyTicket,
		arg.WorkspaceID,
		arg.WorkflowID,
		arg.ID,
		arg.Code,
		arg.Title,
		arg.Priority,
		arg.IsPrivate,
		arg.Status,
		arg.CurrentNodeID,
		arg.CurrentNodeName,
		arg.CurrentNodeStatus,
		arg.CurrentNodeType,
		arg.NodeDeadlineAt,
		arg.AssigneeID,
		arg.DepartmentID,
		arg.Workflow,
		arg.Tags,
		arg.RefTicketIds,
		arg.RatingPoint,
		arg.Review,
		arg.LatestRecentActivityUserID,
		arg.LatestRecentActivityAt,
		arg.CreatedBy,
		arg.UpdatedBy,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const updateAssigneeOfCurrentNodeOnMyTicket = `-- name: UpdateAssigneeOfCurrentNodeOnMyTicket :execresult
UPDATE my_tickets
SET
    assignee_id = ?,
    current_node_status = ?,
    node_deadline_at = ?,
    status = ?,
    updated_by = ?,
    latest_recent_activity_at = ?,
    latest_recent_activity_user_id = ?
WHERE workspace_id = ?
  AND id = ?
  AND current_node_id = ?
`

type UpdateAssigneeOfCurrentNodeOnMyTicketParams struct {
	AssigneeID                 sql.NullInt64
	CurrentNodeStatus          int8
	NodeDeadlineAt             sql.NullTime
	Status                     int8
	UpdatedBy                  int64
	LatestRecentActivityAt     sql.NullTime
	LatestRecentActivityUserID sql.NullInt64
	WorkspaceID                string
	ID                         int64
	CurrentNodeID              int64
}

func (q *Queries) UpdateAssigneeOfCurrentNodeOnMyTicket(ctx context.Context, arg UpdateAssigneeOfCurrentNodeOnMyTicketParams) (sql.Result, error) {
	return q.exec(ctx, q.updateAssigneeOfCurrentNodeOnMyTicketStmt, updateAssigneeOfCurrentNodeOnMyTicket,
		arg.AssigneeID,
		arg.CurrentNodeStatus,
		arg.NodeDeadlineAt,
		arg.Status,
		arg.UpdatedBy,
		arg.LatestRecentActivityAt,
		arg.LatestRecentActivityUserID,
		arg.WorkspaceID,
		arg.ID,
		arg.CurrentNodeID,
	)
}

const updateStatusOfCurrentNodeOnMyTicket = `-- name: UpdateStatusOfCurrentNodeOnMyTicket :execresult
UPDATE my_tickets
SET
    current_node_status = ?,
    status = ?,
    updated_by = ?,
    latest_recent_activity_at = ?,
    latest_recent_activity_user_id = ?
WHERE workspace_id = ?
  AND id = ?
  AND current_node_id = ?
`

type UpdateStatusOfCurrentNodeOnMyTicketParams struct {
	CurrentNodeStatus          int8
	Status                     int8
	UpdatedBy                  int64
	LatestRecentActivityAt     sql.NullTime
	LatestRecentActivityUserID sql.NullInt64
	WorkspaceID                string
	ID                         int64
	CurrentNodeID              int64
}

func (q *Queries) UpdateStatusOfCurrentNodeOnMyTicket(ctx context.Context, arg UpdateStatusOfCurrentNodeOnMyTicketParams) (sql.Result, error) {
	return q.exec(ctx, q.updateStatusOfCurrentNodeOnMyTicketStmt, updateStatusOfCurrentNodeOnMyTicket,
		arg.CurrentNodeStatus,
		arg.Status,
		arg.UpdatedBy,
		arg.LatestRecentActivityAt,
		arg.LatestRecentActivityUserID,
		arg.WorkspaceID,
		arg.ID,
		arg.CurrentNodeID,
	)
}

const updateStatusOfLatestRecentNodeOnMyTicket = `-- name: UpdateStatusOfLatestRecentNodeOnMyTicket :execresult
UPDATE my_tickets
SET status                         = ?,
    updated_by                     = ?,
    updated_at                     = ?,
    latest_recent_activity_at      = ?,
    latest_recent_activity_user_id = ?
WHERE workspace_id = ?
  AND id = ?
  ORDER BY created_at DESC
LIMIT 1
`

type UpdateStatusOfLatestRecentNodeOnMyTicketParams struct {
	Status                     int8
	UpdatedBy                  int64
	UpdatedAt                  time.Time
	LatestRecentActivityAt     sql.NullTime
	LatestRecentActivityUserID sql.NullInt64
	WorkspaceID                string
	ID                         int64
}

func (q *Queries) UpdateStatusOfLatestRecentNodeOnMyTicket(ctx context.Context, arg UpdateStatusOfLatestRecentNodeOnMyTicketParams) (sql.Result, error) {
	return q.exec(ctx, q.updateStatusOfLatestRecentNodeOnMyTicketStmt, updateStatusOfLatestRecentNodeOnMyTicket,
		arg.Status,
		arg.UpdatedBy,
		arg.UpdatedAt,
		arg.LatestRecentActivityAt,
		arg.LatestRecentActivityUserID,
		arg.WorkspaceID,
		arg.ID,
	)
}
