-- name: GetTagWorkflowGroups :many
SELECT * FROM tag_workflow_groups AS t JOIN workflow_groups AS wfg ON t.workflow_group_id=wfg.id
WHERE t.workspace_id=? AND t.tag_id=?;
