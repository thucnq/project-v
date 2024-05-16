-- name: CreateNode :execresult
INSERT INTO nodes
(workspace_id, workflow_id, id, `type`, compressed_option, sla_id, created_at, updated_at)
VALUES(?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateNode :execresult
UPDATE nodes
SET `type` = ?, compressed_option = ?, sla_id = ?, updated_at = ?
WHERE workspace_id = ?
  AND workflow_id = ?
  AND id = ?;

-- name: DeleteAllWorkflowNodes :execresult
DELETE FROM nodes
WHERE workspace_id = ?
  AND workflow_id = ?;

-- name: DeleteNode :execresult
DELETE FROM nodes
WHERE workspace_id = ?
  AND workflow_id = ?
  AND id = ?;

-- name: UpdateSlaIDOfNodes :execresult
UPDATE nodes
SET sla_id = ?
WHERE workspace_id = ?
  AND sla_id = ?;