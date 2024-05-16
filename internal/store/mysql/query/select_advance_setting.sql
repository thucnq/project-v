-- name: GetAdvanceSetting :one
SELECT *
FROM advance_settings
WHERE workspace_id=? AND workflow_id=? LIMIT 1;
