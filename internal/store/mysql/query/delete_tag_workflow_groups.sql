-- name: DeleteTagWorkflowGroups :execresult
DELETE FROM tag_workflow_groups WHERE workspace_id=? AND tag_id=? AND workflow_group_id=?;

-- name: DeleteAllTagWorkflowGroups :execresult
DELETE FROM tag_workflow_groups WHERE workspace_id=? AND tag_id=?;
