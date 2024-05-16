// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: read_ref_my_tickets.sql

package sqlc

import (
	"context"
	"database/sql"
)

const countRefMyTickets = `-- name: CountRefMyTickets :one
SELECT COUNT(DISTINCT id)
FROM (SELECT DISTINCT mt.workspace_id, mt.workflow_id, mt.id, mt.code, mt.title, mt.priority, mt.is_private, mt.status, mt.current_node_id, mt.current_node_name, mt.current_node_status, mt.current_node_type, mt.node_deadline_at, mt.assignee_id, mt.department_id, mt.workflow, mt.tags, mt.ref_ticket_ids, mt.rating_point, mt.review, mt.latest_recent_activity_user_id, mt.latest_recent_activity_at, mt.created_by, mt.updated_by, mt.created_at, mt.updated_at
      FROM my_tickets AS mt
      WHERE mt.workspace_id = ?
        AND mt.assignee_id = ?

      UNION

      SELECT DISTINCT mt.workspace_id, mt.workflow_id, mt.id, mt.code, mt.title, mt.priority, mt.is_private, mt.status, mt.current_node_id, mt.current_node_name, mt.current_node_status, mt.current_node_type, mt.node_deadline_at, mt.assignee_id, mt.department_id, mt.workflow, mt.tags, mt.ref_ticket_ids, mt.rating_point, mt.review, mt.latest_recent_activity_user_id, mt.latest_recent_activity_at, mt.created_by, mt.updated_by, mt.created_at, mt.updated_at
      FROM my_tickets AS mt
               JOIN ticket_nodes AS tn
                    ON mt.id = tn.ticket_id
                        AND mt.current_node_id = tn.node_id
                        AND mt.assignee_id IS NULL
               LEFT JOIN ticket_node_participants AS tnp
                         ON tn.ticket_id = tnp.ticket_id
                             AND tn.node_id = tnp.node_id
               LEFT JOIN participant_members AS pm
                         ON tnp.participant_id = pm.participant_id
                             AND tnp.role_type = 4
      WHERE mt.workspace_id = ?
        AND mt.assignee_id IS NULL
        AND pm.user_id = ?

      UNION

      SELECT DISTINCT mt.workspace_id, mt.workflow_id, mt.id, mt.code, mt.title, mt.priority, mt.is_private, mt.status, mt.current_node_id, mt.current_node_name, mt.current_node_status, mt.current_node_type, mt.node_deadline_at, mt.assignee_id, mt.department_id, mt.workflow, mt.tags, mt.ref_ticket_ids, mt.rating_point, mt.review, mt.latest_recent_activity_user_id, mt.latest_recent_activity_at, mt.created_by, mt.updated_by, mt.created_at, mt.updated_at
      FROM my_tickets AS mt
      WHERE mt.workspace_id = ?
        AND mt.created_by = ?
) AS mt
WHERE ((? = '') OR (mt.title LIKE ?))
`

type CountRefMyTicketsParams struct {
	WorkspaceID   string
	AssigneeID    sql.NullInt64
	WorkspaceID_2 string
	UserID        int64
	WorkspaceID_3 string
	CreatedBy     int64
	Column7       interface{}
	Title         string
}

func (q *Queries) CountRefMyTickets(ctx context.Context, arg CountRefMyTicketsParams) (int64, error) {
	row := q.queryRow(ctx, q.countRefMyTicketsStmt, countRefMyTickets,
		arg.WorkspaceID,
		arg.AssigneeID,
		arg.WorkspaceID_2,
		arg.UserID,
		arg.WorkspaceID_3,
		arg.CreatedBy,
		arg.Column7,
		arg.Title,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getAllRefMyTickets = `-- name: GetAllRefMyTickets :many
SELECT workspace_id, workflow_id, id, code, title, priority, is_private, status, current_node_id, current_node_name, current_node_status, current_node_type, node_deadline_at, assignee_id, department_id, workflow, tags, ref_ticket_ids, rating_point, review, latest_recent_activity_user_id, latest_recent_activity_at, created_by, updated_by, created_at, updated_at
FROM (SELECT DISTINCT mt.workspace_id, mt.workflow_id, mt.id, mt.code, mt.title, mt.priority, mt.is_private, mt.status, mt.current_node_id, mt.current_node_name, mt.current_node_status, mt.current_node_type, mt.node_deadline_at, mt.assignee_id, mt.department_id, mt.workflow, mt.tags, mt.ref_ticket_ids, mt.rating_point, mt.review, mt.latest_recent_activity_user_id, mt.latest_recent_activity_at, mt.created_by, mt.updated_by, mt.created_at, mt.updated_at
      FROM my_tickets AS mt
      WHERE mt.workspace_id = ?
        AND mt.assignee_id = ?

      UNION

      SELECT DISTINCT mt.workspace_id, mt.workflow_id, mt.id, mt.code, mt.title, mt.priority, mt.is_private, mt.status, mt.current_node_id, mt.current_node_name, mt.current_node_status, mt.current_node_type, mt.node_deadline_at, mt.assignee_id, mt.department_id, mt.workflow, mt.tags, mt.ref_ticket_ids, mt.rating_point, mt.review, mt.latest_recent_activity_user_id, mt.latest_recent_activity_at, mt.created_by, mt.updated_by, mt.created_at, mt.updated_at
      FROM my_tickets AS mt
               JOIN ticket_nodes AS tn
                    ON mt.id = tn.ticket_id
                        AND mt.current_node_id = tn.node_id
                        AND mt.assignee_id IS NULL
               LEFT JOIN ticket_node_participants AS tnp
                         ON tn.ticket_id = tnp.ticket_id
                             AND tn.node_id = tnp.node_id
               LEFT JOIN participant_members AS pm
                         ON tnp.participant_id = pm.participant_id
                             AND tnp.role_type = 4
      WHERE mt.workspace_id = ?
        AND mt.assignee_id IS NULL
        AND pm.user_id = ?

      UNION

      SELECT DISTINCT mt.workspace_id, mt.workflow_id, mt.id, mt.code, mt.title, mt.priority, mt.is_private, mt.status, mt.current_node_id, mt.current_node_name, mt.current_node_status, mt.current_node_type, mt.node_deadline_at, mt.assignee_id, mt.department_id, mt.workflow, mt.tags, mt.ref_ticket_ids, mt.rating_point, mt.review, mt.latest_recent_activity_user_id, mt.latest_recent_activity_at, mt.created_by, mt.updated_by, mt.created_at, mt.updated_at
      FROM my_tickets AS mt
      WHERE mt.workspace_id = ?
            AND mt.created_by = ?
) AS mt
WHERE ((? = '') OR (mt.title LIKE ?))
`

type GetAllRefMyTicketsParams struct {
	WorkspaceID   string
	AssigneeID    sql.NullInt64
	WorkspaceID_2 string
	UserID        int64
	WorkspaceID_3 string
	CreatedBy     int64
	Column7       interface{}
	Title         string
}

func (q *Queries) GetAllRefMyTickets(ctx context.Context, arg GetAllRefMyTicketsParams) ([]MyTicket, error) {
	rows, err := q.query(ctx, q.getAllRefMyTicketsStmt, getAllRefMyTickets,
		arg.WorkspaceID,
		arg.AssigneeID,
		arg.WorkspaceID_2,
		arg.UserID,
		arg.WorkspaceID_3,
		arg.CreatedBy,
		arg.Column7,
		arg.Title,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MyTicket
	for rows.Next() {
		var i MyTicket
		if err := rows.Scan(
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
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
