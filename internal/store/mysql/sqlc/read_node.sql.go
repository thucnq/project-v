// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: read_node.sql

package sqlc

import (
	"context"
)

const countNodesByWorkflowID = `-- name: CountNodesByWorkflowID :one
SELECT COUNT(*)
FROM nodes
WHERE workspace_id = ? AND workflow_id = ?
`

type CountNodesByWorkflowIDParams struct {
	WorkspaceID string
	WorkflowID  int64
}

func (q *Queries) CountNodesByWorkflowID(ctx context.Context, arg CountNodesByWorkflowIDParams) (int64, error) {
	row := q.queryRow(ctx, q.countNodesByWorkflowIDStmt, countNodesByWorkflowID, arg.WorkspaceID, arg.WorkflowID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getNodeByID = `-- name: GetNodeByID :one
SELECT workspace_id, workflow_id, id, type, compressed_option, sla_id, created_at, updated_at
FROM nodes
WHERE workspace_id = ? AND workflow_id = ? AND id = ?
`

type GetNodeByIDParams struct {
	WorkspaceID string
	WorkflowID  int64
	ID          int64
}

func (q *Queries) GetNodeByID(ctx context.Context, arg GetNodeByIDParams) (Node, error) {
	row := q.queryRow(ctx, q.getNodeByIDStmt, getNodeByID, arg.WorkspaceID, arg.WorkflowID, arg.ID)
	var i Node
	err := row.Scan(
		&i.WorkspaceID,
		&i.WorkflowID,
		&i.ID,
		&i.Type,
		&i.CompressedOption,
		&i.SlaID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getStartNode = `-- name: GetStartNode :one
SELECT workspace_id, workflow_id, id, type, compressed_option, sla_id, created_at, updated_at
FROM nodes
WHERE workspace_id = ? AND workflow_id = ? AND type = 1
`

type GetStartNodeParams struct {
	WorkspaceID string
	WorkflowID  int64
}

func (q *Queries) GetStartNode(ctx context.Context, arg GetStartNodeParams) (Node, error) {
	row := q.queryRow(ctx, q.getStartNodeStmt, getStartNode, arg.WorkspaceID, arg.WorkflowID)
	var i Node
	err := row.Scan(
		&i.WorkspaceID,
		&i.WorkflowID,
		&i.ID,
		&i.Type,
		&i.CompressedOption,
		&i.SlaID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listNodesByWorkflowID = `-- name: ListNodesByWorkflowID :many
SELECT workspace_id, workflow_id, id, type, compressed_option, sla_id, created_at, updated_at
FROM nodes
WHERE workspace_id = ? AND workflow_id = ?
`

type ListNodesByWorkflowIDParams struct {
	WorkspaceID string
	WorkflowID  int64
}

func (q *Queries) ListNodesByWorkflowID(ctx context.Context, arg ListNodesByWorkflowIDParams) ([]Node, error) {
	rows, err := q.query(ctx, q.listNodesByWorkflowIDStmt, listNodesByWorkflowID, arg.WorkspaceID, arg.WorkflowID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Node
	for rows.Next() {
		var i Node
		if err := rows.Scan(
			&i.WorkspaceID,
			&i.WorkflowID,
			&i.ID,
			&i.Type,
			&i.CompressedOption,
			&i.SlaID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
