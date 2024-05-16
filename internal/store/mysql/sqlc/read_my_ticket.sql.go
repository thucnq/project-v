// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: read_my_ticket.sql

package sqlc

import (
	"context"
)

const countTotalNodeOnMyTicket = `-- name: CountTotalNodeOnMyTicket :one
SELECT COUNT(*)
FROM my_tickets
WHERE workspace_id = ?
  AND id = ?
`

type CountTotalNodeOnMyTicketParams struct {
	WorkspaceID string
	ID          int64
}

func (q *Queries) CountTotalNodeOnMyTicket(ctx context.Context, arg CountTotalNodeOnMyTicketParams) (int64, error) {
	row := q.queryRow(ctx, q.countTotalNodeOnMyTicketStmt, countTotalNodeOnMyTicket, arg.WorkspaceID, arg.ID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getMyTicketByCurrentNodeID = `-- name: GetMyTicketByCurrentNodeID :one
SELECT workspace_id, workflow_id, id, code, title, priority, is_private, status, current_node_id, current_node_name, current_node_status, current_node_type, node_deadline_at, assignee_id, department_id, workflow, tags, ref_ticket_ids, rating_point, review, latest_recent_activity_user_id, latest_recent_activity_at, created_by, updated_by, created_at, updated_at
FROM my_tickets
WHERE workspace_id = ?
  AND id = ?
  AND current_node_id = ?
`

type GetMyTicketByCurrentNodeIDParams struct {
	WorkspaceID   string
	ID            int64
	CurrentNodeID int64
}

func (q *Queries) GetMyTicketByCurrentNodeID(ctx context.Context, arg GetMyTicketByCurrentNodeIDParams) (MyTicket, error) {
	row := q.queryRow(ctx, q.getMyTicketByCurrentNodeIDStmt, getMyTicketByCurrentNodeID, arg.WorkspaceID, arg.ID, arg.CurrentNodeID)
	var i MyTicket
	err := row.Scan(
		&i.WorkspaceID,
		&i.WorkflowID,
		&i.ID,
		&i.Code,
		&i.Title,
		&i.Priority,
		&i.IsPrivate,
		&i.Status,
		&i.CurrentNodeID,
		&i.CurrentNodeName,
		&i.CurrentNodeStatus,
		&i.CurrentNodeType,
		&i.NodeDeadlineAt,
		&i.AssigneeID,
		&i.DepartmentID,
		&i.Workflow,
		&i.Tags,
		&i.RefTicketIds,
		&i.RatingPoint,
		&i.Review,
		&i.LatestRecentActivityUserID,
		&i.LatestRecentActivityAt,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
