-- name: CreateWorkflowGroup :execresult
INSERT INTO workflow_groups
(workspace_id, parent_id, id, name, left_bower, right_bower, h_level)
VALUES(?, ?, ?, ?, ?, ?, ?);

-- name: UpdateBowersBeforeAppend :execresult
CALL UpdateBowersBeforeAppend(?,?);

-- name: UpdateWorkflowGroup :execresult
UPDATE workflow_groups SET name=?, updated_at=?, version = version + 1 WHERE workspace_id = ? AND parent_id > 0 AND id = ?;

-- name: DeleteRange :execresult
CALL DeleteRange(?,?,?,?);