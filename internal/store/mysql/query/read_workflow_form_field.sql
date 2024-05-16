-- name: GetWorkflowFormField :one
SELECT *
FROM workflow_form_fields
WHERE workspace_id=? AND workflow_id=? AND id=? LIMIT 1;

-- name: ListWorkflowFormFieldsByWorkflowId :many
SELECT *
FROM workflow_form_fields
WHERE workspace_id=? AND workflow_id=?;

-- name: ListFirstClassWorkflowFormFieldsByWorkflowId :many
SELECT *
FROM workflow_form_fields
WHERE workspace_id=? AND workflow_id=? AND parent_id = 0;


