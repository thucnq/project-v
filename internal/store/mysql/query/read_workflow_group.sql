-- name: GetWorkflowGroup :one
SELECT *
FROM workflow_groups
WHERE workspace_id=? AND id=? LIMIT 1;

-- name: GetRootWorkflowGroup :one
SELECT *
FROM workflow_groups
WHERE workspace_id=? AND parent_id=0 LIMIT 1;

-- name: GetRootWorkflowGroupsForUpdate :many
SELECT *
FROM workflow_groups
WHERE workspace_id = ? FOR UPDATE;

-- name: GetChildWorkflowGroups :many
SELECT *
FROM workflow_groups
WHERE workspace_id=? AND parent_id=?
ORDER BY left_bower ASC;

-- name: GetAllWorkflowGroups :many
SELECT *
FROM workflow_groups
WHERE workspace_id=? AND parent_id > 0
ORDER BY left_bower ASC;

-- name: GetRangeWorkflowGroups :many
SELECT *
FROM workflow_groups
WHERE workspace_id=? AND left_bower > ? AND right_bower < ?
ORDER BY left_bower ASC;

-- name: GetWorkflowGroups :many
SELECT *
FROM workflow_groups
WHERE workspace_id=?;

-- name: GetNextWorkflowGroupsByKeyword :many
SELECT *
FROM workflow_groups
WHERE workspace_id=? AND id < ? AND name LIKE ?
ORDER BY id DESC
LIMIT ?;

-- name: GetPreviousWorkflowGroupsByKeyword :many
SELECT *
FROM workflow_groups
WHERE workspace_id=? AND id > ? AND name LIKE ?
ORDER BY id ASC
LIMIT ?;

-- name: GetWorkflowGroupsByKeyword :many
SELECT *
FROM workflow_groups
WHERE workspace_id=? AND name LIKE ?
ORDER BY id DESC
LIMIT ?;

-- name: GetFullPathOfWorkflowGroup :one
SELECT GROUP_CONCAT(p.name SEPARATOR '/') FROM(
        SELECT parent.name
        FROM workflow_groups AS node,
             workflow_groups AS parent
        WHERE parent.workspace_id = ?
          AND node.id = ?
          AND node.left_bower
              BETWEEN parent.left_bower AND parent.right_bower
        ORDER BY parent.left_bower ASC
        LIMIT 10
        OFFSET 1
) AS p;

