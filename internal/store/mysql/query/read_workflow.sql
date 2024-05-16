-- name: GetWorkflowByID :one
SELECT
    workflows.*, workflow_groups.name AS workflow_group_name
FROM workflows
INNER JOIN workflow_groups
    ON workflows.workflow_group_id = workflow_groups.id
WHERE workflows.id = ?;

-- name: GetAllWorkflows :many
SELECT workflows.*,
       workflow_groups.name AS workflow_group_name
FROM workflows
INNER JOIN workflow_groups
    ON workflows.workflow_group_id = workflow_groups.id
WHERE workflows.workspace_id = ?
  AND ( -- Nếu get_of_group_child là false, chỉ lấy các workflows thuộc group hiện tại
        (workflows.workflow_group_id = ?) OR workflows.workflow_group_id IN (
                                                SELECT id
                                                FROM workflow_groups
                                                WHERE workflow_groups.left_bower >= ?
                                                  AND workflow_groups.right_bower <= ?)
    )
  AND (workflows.state = ? OR workflows.state IS NULL OR ? = 0)
  AND ((? = 0) OR (workflows.is_published = ? AND ? = 1) OR (workflows.is_published = FALSE AND ? = 2))
  AND ((? = '') OR (workflows.name LIKE ?))
ORDER BY workflows.id DESC LIMIT ?;

-- name: CountWorkflows :one
SELECT COUNT(*)
FROM workflows
INNER JOIN workflow_groups
    ON workflows.workflow_group_id = workflow_groups.id
WHERE workflows.workspace_id = ?
  AND ( -- Nếu get_of_group_child là false, chỉ lấy các workflows thuộc group hiện tại
        (workflows.workflow_group_id = ?) OR workflows.workflow_group_id IN (
                                                    SELECT id
                                                    FROM workflow_groups
                                                    WHERE workflow_groups.left_bower >= ?
                                                      AND workflow_groups.right_bower <= ?)
    )
  AND (workflows.state = ? OR workflows.state IS NULL OR ? = 0)
  AND ((? = 0) OR (workflows.is_published = ? AND ? = 1) OR (workflows.is_published = FALSE AND ? = 2))
  AND ((? = '') OR (workflows.name LIKE ?));

-- name: GetAllWorkflowsPagingNext :many
SELECT workflows.*,
       workflow_groups.name AS workflow_group_name
FROM workflows
INNER JOIN workflow_groups
    ON workflows.workflow_group_id = workflow_groups.id
WHERE workflows.workspace_id = ?
  AND ( -- Nếu get_of_group_child là false, chỉ lấy các workflows thuộc group hiện tại
        (workflows.workflow_group_id = ?) OR workflows.workflow_group_id IN (
                                                        SELECT id
                                                        FROM workflow_groups
                                                        WHERE workflow_groups.left_bower >= ?
                                                          AND workflow_groups.right_bower <= ?)
    )
  AND (workflows.state = ? OR workflows.state IS NULL OR ? = 0)
  AND ((? = 0) OR (workflows.is_published = ? AND ? = 1) OR (workflows.is_published = FALSE AND ? = 2))
  AND ((? = '') OR (workflows.name LIKE ?))
  AND workflows.id < ?
ORDER BY workflows.id DESC
LIMIT ?;

-- name: GetAllWorkflowsPagingPrev :many
SELECT workflows.*,
       workflow_groups.name AS workflow_group_name
FROM workflows
INNER JOIN workflow_groups
    ON workflows.workflow_group_id = workflow_groups.id
WHERE workflows.workspace_id = ?
  AND ( -- Nếu get_of_group_child là false, chỉ lấy các workflows thuộc group hiện tại
        (workflows.workflow_group_id = ?) OR workflows.workflow_group_id IN (
                                                        SELECT id
                                                        FROM workflow_groups
                                                        WHERE workflow_groups.left_bower >= ?
                                                          AND workflow_groups.right_bower <= ?)
    )
  AND (workflows.state = ? OR workflows.state IS NULL OR ? = 0)
  AND ((? = 0) OR (workflows.is_published = ? AND ? = 1) OR (workflows.is_published = FALSE AND ? = 2))
  AND ((? = '') OR (workflows.name LIKE ?))
  AND workflows.id > ?
ORDER BY workflows.id ASC
LIMIT ?;