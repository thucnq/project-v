-- name: CreateMyTicket :execresult
INSERT INTO my_tickets (
    workspace_id, workflow_id, id,
    code, title, priority, is_private, status,
    current_node_id, current_node_name, current_node_status, current_node_type,
    node_deadline_at, assignee_id,
    department_id, workflow, tags, ref_ticket_ids,
    rating_point, review,
    latest_recent_activity_user_id, latest_recent_activity_at,
    created_by, updated_by, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateAssigneeOfCurrentNodeOnMyTicket :execresult
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
  AND current_node_id = ?;

-- name: UpdateStatusOfCurrentNodeOnMyTicket :execresult
UPDATE my_tickets
SET
    current_node_status = ?,
    status = ?,
    updated_by = ?,
    latest_recent_activity_at = ?,
    latest_recent_activity_user_id = ?
WHERE workspace_id = ?
  AND id = ?
  AND current_node_id = ?;

-- name: UpdateStatusOfLatestRecentNodeOnMyTicket :execresult
UPDATE my_tickets
SET status                         = ?,
    updated_by                     = ?,
    updated_at                     = ?,
    latest_recent_activity_at      = ?,
    latest_recent_activity_user_id = ?
WHERE workspace_id = ?
  AND id = ?
  ORDER BY created_at DESC
LIMIT 1;