-- name: CreateWorkflowFormField :execresult
INSERT INTO workflow_form_fields
(workspace_id, workflow_id, id, title, hint, type, is_required, compressed_option, parent_id)
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateWorkflowFormField :execresult
UPDATE workflow_form_fields SET title=?, type=?,is_required=?,compressed_option=?, updated_at=?, version = version + 1 WHERE workspace_id = ? AND parent_id > 0 AND id = ?;

-- name: DeleteWorkflowFormFields :execresult
DELETE FROM workflow_form_fields WHERE workspace_id = ? AND workflow_id = ? AND id IN (?);

-- name: DeleteOtherWorkflowFormFields :execresult
DELETE FROM workflow_form_fields WHERE workspace_id = ? AND workflow_id = ? AND id  NOT IN (?);

-- name: UpsertWorkflowFormField :execresult
INSERT INTO workflow_form_fields (workspace_id, workflow_id, id, title, hint, type, is_required, compressed_option, parent_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE title=VALUES(title),type=VALUES(type),is_required=VALUES(is_required),compressed_option=VALUES(compressed_option),parent_id=VALUES(parent_id),updated_at=?, version = version + 1;

-- name: DeleteWorkflowFormFieldsByWorkflowId :execresult
DELETE FROM workflow_form_fields WHERE workspace_id = ? AND workflow_id = ?;

-- name: DeleteWorkflowFormField :execresult
DELETE FROM workflow_form_fields WHERE workspace_id = ? AND workflow_id = ? AND id = ?;