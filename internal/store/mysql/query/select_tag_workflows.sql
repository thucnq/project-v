-- name: GetTagWorkflows :many
SELECT * FROM tag_workflows AS t JOIN workflows AS wf ON t.workflow_id=wf.id
WHERE t.workspace_id=? AND t.tag_id=?;
