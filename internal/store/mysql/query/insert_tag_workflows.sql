-- name: CreateTagWorkflows :execresult
INSERT IGNORE INTO tag_workflows (
	workspace_id,
	tag_id,
	workflow_id
)
VALUES (
	?, ?, ?
);
