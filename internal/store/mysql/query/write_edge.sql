-- name: CreateEdge :execresult
INSERT INTO edges
(workspace_id, workflow_id, current_node_id, next_node_id, created_at, updated_at)
VALUES(?, ?, ?, ?, ?, ?);

-- name: DeleteNodeEdges :execresult
DELETE FROM edges
WHERE workspace_id = ?
  AND workflow_id = ?
  AND (current_node_id = ? OR next_node_id = ?);

-- name: DeleteAllWorkflowEdges :execresult
DELETE FROM edges
WHERE workspace_id = ?
  AND workflow_id = ?;

-- name: DeleteEdges :execresult
DELETE FROM edges
WHERE workspace_id = ?
  AND workflow_id = ?
  AND current_node_id = ? AND next_node_id = ?;