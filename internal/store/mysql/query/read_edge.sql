-- name: ListEdgesByWorkflowID :many
SELECT *
FROM edges
WHERE workspace_id = ? AND workflow_id = ?;

-- name: GetEdgeByID :one
SELECT *
FROM edges
WHERE workspace_id = ? AND workflow_id = ? AND current_node_id = ? and next_node_id;