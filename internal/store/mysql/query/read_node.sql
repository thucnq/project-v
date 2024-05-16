-- name: CountNodesByWorkflowID :one
SELECT COUNT(*)
FROM nodes
WHERE workspace_id = ? AND workflow_id = ?;

-- name: ListNodesByWorkflowID :many
SELECT *
FROM nodes
WHERE workspace_id = ? AND workflow_id = ?;

-- name: GetNodeByID :one
SELECT *
FROM nodes
WHERE workspace_id = ? AND workflow_id = ? AND id = ?;

-- name: GetStartNode :one
SELECT *
FROM nodes
WHERE workspace_id = ? AND workflow_id = ? AND type = 1;