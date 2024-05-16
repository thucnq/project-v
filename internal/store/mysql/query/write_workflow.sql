-- name: CreateWorkflow :execresult
INSERT INTO workflows (
    id, workspace_id, workflow_group_id, prefix_id, name, color_src, icon_src,
    description, state, created_by, updated_by, is_published, created_at, updated_at
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: DeleteWorkflowByID :execresult
DELETE FROM workflows
WHERE workspace_id = ?
    AND id = ?;

-- name: UpdateWorkflowByWorkspaceIDAndID :execresult
UPDATE workflows
SET
    name = ?, color_src = ?, icon_src = ?,
    description = ?, workflow_group_id = ?, prefix_id = ?,
    state = ?, updated_by = ?, is_published = ?, updated_at = ?
WHERE workspace_id = ?
  AND id = ?;

-- name: UpdateWorkflowPublishedByWorkspaceIDAndID :execresult
UPDATE workflows
SET is_published = ?, updated_by = ?, updated_at = ?
WHERE workspace_id = ?
  AND id = ?;

-- name: UpdatePrefixIDOfWorkflows :execresult
UPDATE workflows
SET prefix_id = ?
WHERE workspace_id = ?
  AND prefix_id = ?;