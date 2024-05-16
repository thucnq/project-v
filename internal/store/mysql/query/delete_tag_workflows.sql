-- name: DeleteTagWorkflows :execresult
DELETE FROM tag_workflows WHERE workspace_id=? AND tag_id=? AND workflow_id=?;

-- name: DeleteAllTagWorkflows :execresult
DELETE FROM tag_workflows WHERE workspace_id=? AND tag_id=?;
