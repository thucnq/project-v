-- name: CreateTagWorkflowGroups :execresult
INSERT IGNORE INTO tag_workflow_groups (
	workspace_id,
	tag_id,
	workflow_group_id
)
VALUES (
	?, ?, ?
);
