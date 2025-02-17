// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: delete_tag_workflow_groups.sql

package sqlc

import (
	"context"
	"database/sql"
)

const deleteAllTagWorkflowGroups = `-- name: DeleteAllTagWorkflowGroups :execresult
DELETE FROM tag_workflow_groups WHERE workspace_id=? AND tag_id=?
`

type DeleteAllTagWorkflowGroupsParams struct {
	WorkspaceID string
	TagID       int64
}

func (q *Queries) DeleteAllTagWorkflowGroups(ctx context.Context, arg DeleteAllTagWorkflowGroupsParams) (sql.Result, error) {
	return q.exec(ctx, q.deleteAllTagWorkflowGroupsStmt, deleteAllTagWorkflowGroups, arg.WorkspaceID, arg.TagID)
}

const deleteTagWorkflowGroups = `-- name: DeleteTagWorkflowGroups :execresult
DELETE FROM tag_workflow_groups WHERE workspace_id=? AND tag_id=? AND workflow_group_id=?
`

type DeleteTagWorkflowGroupsParams struct {
	WorkspaceID     string
	TagID           int64
	WorkflowGroupID int64
}

func (q *Queries) DeleteTagWorkflowGroups(ctx context.Context, arg DeleteTagWorkflowGroupsParams) (sql.Result, error) {
	return q.exec(ctx, q.deleteTagWorkflowGroupsStmt, deleteTagWorkflowGroups, arg.WorkspaceID, arg.TagID, arg.WorkflowGroupID)
}
