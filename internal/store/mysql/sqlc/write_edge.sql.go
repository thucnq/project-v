// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: write_edge.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createEdge = `-- name: CreateEdge :execresult
INSERT INTO edges
(workspace_id, workflow_id, current_node_id, next_node_id, created_at, updated_at)
VALUES(?, ?, ?, ?, ?, ?)
`

type CreateEdgeParams struct {
	WorkspaceID   string
	WorkflowID    int64
	CurrentNodeID int64
	NextNodeID    int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (q *Queries) CreateEdge(ctx context.Context, arg CreateEdgeParams) (sql.Result, error) {
	return q.exec(ctx, q.createEdgeStmt, createEdge,
		arg.WorkspaceID,
		arg.WorkflowID,
		arg.CurrentNodeID,
		arg.NextNodeID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const deleteAllWorkflowEdges = `-- name: DeleteAllWorkflowEdges :execresult
DELETE FROM edges
WHERE workspace_id = ?
  AND workflow_id = ?
`

type DeleteAllWorkflowEdgesParams struct {
	WorkspaceID string
	WorkflowID  int64
}

func (q *Queries) DeleteAllWorkflowEdges(ctx context.Context, arg DeleteAllWorkflowEdgesParams) (sql.Result, error) {
	return q.exec(ctx, q.deleteAllWorkflowEdgesStmt, deleteAllWorkflowEdges, arg.WorkspaceID, arg.WorkflowID)
}

const deleteEdges = `-- name: DeleteEdges :execresult
DELETE FROM edges
WHERE workspace_id = ?
  AND workflow_id = ?
  AND current_node_id = ? AND next_node_id = ?
`

type DeleteEdgesParams struct {
	WorkspaceID   string
	WorkflowID    int64
	CurrentNodeID int64
	NextNodeID    int64
}

func (q *Queries) DeleteEdges(ctx context.Context, arg DeleteEdgesParams) (sql.Result, error) {
	return q.exec(ctx, q.deleteEdgesStmt, deleteEdges,
		arg.WorkspaceID,
		arg.WorkflowID,
		arg.CurrentNodeID,
		arg.NextNodeID,
	)
}

const deleteNodeEdges = `-- name: DeleteNodeEdges :execresult
DELETE FROM edges
WHERE workspace_id = ?
  AND workflow_id = ?
  AND (current_node_id = ? OR next_node_id = ?)
`

type DeleteNodeEdgesParams struct {
	WorkspaceID   string
	WorkflowID    int64
	CurrentNodeID int64
	NextNodeID    int64
}

func (q *Queries) DeleteNodeEdges(ctx context.Context, arg DeleteNodeEdgesParams) (sql.Result, error) {
	return q.exec(ctx, q.deleteNodeEdgesStmt, deleteNodeEdges,
		arg.WorkspaceID,
		arg.WorkflowID,
		arg.CurrentNodeID,
		arg.NextNodeID,
	)
}
