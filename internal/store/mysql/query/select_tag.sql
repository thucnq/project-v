-- name: GetTag :one
SELECT *
FROM tags
WHERE workspace_id=? AND id=? LIMIT 1;

-- name: GetTagsSortByName :many
SELECT *
FROM tags
WHERE workspace_id=? ORDER BY created_at DESC LIMIT ? OFFSET ?;

-- name: GetTagSearchByName :many
SELECT *
FROM tags
WHERE workspace_id=? AND name LIKE ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetTagSearchByNameWithPublished :many
SELECT *
FROM tags
WHERE workspace_id=? AND name LIKE ? AND is_published=?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetTagWithPublished :many
SELECT *
FROM tags
WHERE workspace_id=? AND is_published=?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetTagsByWorkflowID :many
SELECT t.*
FROM tags AS t
    JOIN tag_workflows AS tw
        ON tw.workspace_id = t.workspace_id
        AND tw.tag_id = t.id
WHERE tw.workspace_id = ?
  AND (tw.workflow_id = ? OR tw.workflow_id = 0)
  AND t.is_deleted = false
  AND t.is_published = true;