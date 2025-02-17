// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: read_task_participant.sql

package sqlc

import (
	"context"
)

const listTaskParticipantsByEightTaskIDs = `-- name: ListTaskParticipantsByEightTaskIDs :many
SELECT workspace_id, task_id, participant_id, role_type, participant_type, created_at, updated_at
FROM task_participants
WHERE  workspace_id = ? AND task_id IN (?,?,?,?,?,?,?,?)
`

type ListTaskParticipantsByEightTaskIDsParams struct {
	WorkspaceID string
	TaskID      int64
	TaskID_2    int64
	TaskID_3    int64
	TaskID_4    int64
	TaskID_5    int64
	TaskID_6    int64
	TaskID_7    int64
	TaskID_8    int64
}

func (q *Queries) ListTaskParticipantsByEightTaskIDs(ctx context.Context, arg ListTaskParticipantsByEightTaskIDsParams) ([]TaskParticipant, error) {
	rows, err := q.query(ctx, q.listTaskParticipantsByEightTaskIDsStmt, listTaskParticipantsByEightTaskIDs,
		arg.WorkspaceID,
		arg.TaskID,
		arg.TaskID_2,
		arg.TaskID_3,
		arg.TaskID_4,
		arg.TaskID_5,
		arg.TaskID_6,
		arg.TaskID_7,
		arg.TaskID_8,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TaskParticipant
	for rows.Next() {
		var i TaskParticipant
		if err := rows.Scan(
			&i.WorkspaceID,
			&i.TaskID,
			&i.ParticipantID,
			&i.RoleType,
			&i.ParticipantType,
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

const listTaskParticipantsByFiveTaskIDs = `-- name: ListTaskParticipantsByFiveTaskIDs :many
SELECT workspace_id, task_id, participant_id, role_type, participant_type, created_at, updated_at
FROM task_participants
WHERE  workspace_id = ? AND task_id IN (?,?,?,?,?)
`

type ListTaskParticipantsByFiveTaskIDsParams struct {
	WorkspaceID string
	TaskID      int64
	TaskID_2    int64
	TaskID_3    int64
	TaskID_4    int64
	TaskID_5    int64
}

func (q *Queries) ListTaskParticipantsByFiveTaskIDs(ctx context.Context, arg ListTaskParticipantsByFiveTaskIDsParams) ([]TaskParticipant, error) {
	rows, err := q.query(ctx, q.listTaskParticipantsByFiveTaskIDsStmt, listTaskParticipantsByFiveTaskIDs,
		arg.WorkspaceID,
		arg.TaskID,
		arg.TaskID_2,
		arg.TaskID_3,
		arg.TaskID_4,
		arg.TaskID_5,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TaskParticipant
	for rows.Next() {
		var i TaskParticipant
		if err := rows.Scan(
			&i.WorkspaceID,
			&i.TaskID,
			&i.ParticipantID,
			&i.RoleType,
			&i.ParticipantType,
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

const listTaskParticipantsByFourTaskIDs = `-- name: ListTaskParticipantsByFourTaskIDs :many
SELECT workspace_id, task_id, participant_id, role_type, participant_type, created_at, updated_at
FROM task_participants
WHERE  workspace_id = ? AND task_id IN (?,?,?,?)
`

type ListTaskParticipantsByFourTaskIDsParams struct {
	WorkspaceID string
	TaskID      int64
	TaskID_2    int64
	TaskID_3    int64
	TaskID_4    int64
}

func (q *Queries) ListTaskParticipantsByFourTaskIDs(ctx context.Context, arg ListTaskParticipantsByFourTaskIDsParams) ([]TaskParticipant, error) {
	rows, err := q.query(ctx, q.listTaskParticipantsByFourTaskIDsStmt, listTaskParticipantsByFourTaskIDs,
		arg.WorkspaceID,
		arg.TaskID,
		arg.TaskID_2,
		arg.TaskID_3,
		arg.TaskID_4,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TaskParticipant
	for rows.Next() {
		var i TaskParticipant
		if err := rows.Scan(
			&i.WorkspaceID,
			&i.TaskID,
			&i.ParticipantID,
			&i.RoleType,
			&i.ParticipantType,
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

const listTaskParticipantsByNineTaskIDs = `-- name: ListTaskParticipantsByNineTaskIDs :many
SELECT workspace_id, task_id, participant_id, role_type, participant_type, created_at, updated_at
FROM task_participants
WHERE  workspace_id = ? AND task_id IN (?,?,?,?,?,?,?,?,?)
`

type ListTaskParticipantsByNineTaskIDsParams struct {
	WorkspaceID string
	TaskID      int64
	TaskID_2    int64
	TaskID_3    int64
	TaskID_4    int64
	TaskID_5    int64
	TaskID_6    int64
	TaskID_7    int64
	TaskID_8    int64
	TaskID_9    int64
}

func (q *Queries) ListTaskParticipantsByNineTaskIDs(ctx context.Context, arg ListTaskParticipantsByNineTaskIDsParams) ([]TaskParticipant, error) {
	rows, err := q.query(ctx, q.listTaskParticipantsByNineTaskIDsStmt, listTaskParticipantsByNineTaskIDs,
		arg.WorkspaceID,
		arg.TaskID,
		arg.TaskID_2,
		arg.TaskID_3,
		arg.TaskID_4,
		arg.TaskID_5,
		arg.TaskID_6,
		arg.TaskID_7,
		arg.TaskID_8,
		arg.TaskID_9,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TaskParticipant
	for rows.Next() {
		var i TaskParticipant
		if err := rows.Scan(
			&i.WorkspaceID,
			&i.TaskID,
			&i.ParticipantID,
			&i.RoleType,
			&i.ParticipantType,
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

const listTaskParticipantsBySevenTaskIDs = `-- name: ListTaskParticipantsBySevenTaskIDs :many
SELECT workspace_id, task_id, participant_id, role_type, participant_type, created_at, updated_at
FROM task_participants
WHERE  workspace_id = ? AND task_id IN (?,?,?,?,?,?,?)
`

type ListTaskParticipantsBySevenTaskIDsParams struct {
	WorkspaceID string
	TaskID      int64
	TaskID_2    int64
	TaskID_3    int64
	TaskID_4    int64
	TaskID_5    int64
	TaskID_6    int64
	TaskID_7    int64
}

func (q *Queries) ListTaskParticipantsBySevenTaskIDs(ctx context.Context, arg ListTaskParticipantsBySevenTaskIDsParams) ([]TaskParticipant, error) {
	rows, err := q.query(ctx, q.listTaskParticipantsBySevenTaskIDsStmt, listTaskParticipantsBySevenTaskIDs,
		arg.WorkspaceID,
		arg.TaskID,
		arg.TaskID_2,
		arg.TaskID_3,
		arg.TaskID_4,
		arg.TaskID_5,
		arg.TaskID_6,
		arg.TaskID_7,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TaskParticipant
	for rows.Next() {
		var i TaskParticipant
		if err := rows.Scan(
			&i.WorkspaceID,
			&i.TaskID,
			&i.ParticipantID,
			&i.RoleType,
			&i.ParticipantType,
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

const listTaskParticipantsBySixTaskIDs = `-- name: ListTaskParticipantsBySixTaskIDs :many
SELECT workspace_id, task_id, participant_id, role_type, participant_type, created_at, updated_at
FROM task_participants
WHERE  workspace_id = ? AND task_id IN (?,?,?,?,?,?)
`

type ListTaskParticipantsBySixTaskIDsParams struct {
	WorkspaceID string
	TaskID      int64
	TaskID_2    int64
	TaskID_3    int64
	TaskID_4    int64
	TaskID_5    int64
	TaskID_6    int64
}

func (q *Queries) ListTaskParticipantsBySixTaskIDs(ctx context.Context, arg ListTaskParticipantsBySixTaskIDsParams) ([]TaskParticipant, error) {
	rows, err := q.query(ctx, q.listTaskParticipantsBySixTaskIDsStmt, listTaskParticipantsBySixTaskIDs,
		arg.WorkspaceID,
		arg.TaskID,
		arg.TaskID_2,
		arg.TaskID_3,
		arg.TaskID_4,
		arg.TaskID_5,
		arg.TaskID_6,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TaskParticipant
	for rows.Next() {
		var i TaskParticipant
		if err := rows.Scan(
			&i.WorkspaceID,
			&i.TaskID,
			&i.ParticipantID,
			&i.RoleType,
			&i.ParticipantType,
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

const listTaskParticipantsByTaskID = `-- name: ListTaskParticipantsByTaskID :many
SELECT workspace_id, task_id, participant_id, role_type, participant_type, created_at, updated_at
FROM task_participants
WHERE workspace_id = ? AND task_id = ?
`

type ListTaskParticipantsByTaskIDParams struct {
	WorkspaceID string
	TaskID      int64
}

func (q *Queries) ListTaskParticipantsByTaskID(ctx context.Context, arg ListTaskParticipantsByTaskIDParams) ([]TaskParticipant, error) {
	rows, err := q.query(ctx, q.listTaskParticipantsByTaskIDStmt, listTaskParticipantsByTaskID, arg.WorkspaceID, arg.TaskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TaskParticipant
	for rows.Next() {
		var i TaskParticipant
		if err := rows.Scan(
			&i.WorkspaceID,
			&i.TaskID,
			&i.ParticipantID,
			&i.RoleType,
			&i.ParticipantType,
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

const listTaskParticipantsByTenTaskIDs = `-- name: ListTaskParticipantsByTenTaskIDs :many
SELECT workspace_id, task_id, participant_id, role_type, participant_type, created_at, updated_at
FROM task_participants
WHERE  workspace_id = ? AND task_id IN (?,?,?,?,?,?,?,?,?,?)
`

type ListTaskParticipantsByTenTaskIDsParams struct {
	WorkspaceID string
	TaskID      int64
	TaskID_2    int64
	TaskID_3    int64
	TaskID_4    int64
	TaskID_5    int64
	TaskID_6    int64
	TaskID_7    int64
	TaskID_8    int64
	TaskID_9    int64
	TaskID_10   int64
}

func (q *Queries) ListTaskParticipantsByTenTaskIDs(ctx context.Context, arg ListTaskParticipantsByTenTaskIDsParams) ([]TaskParticipant, error) {
	rows, err := q.query(ctx, q.listTaskParticipantsByTenTaskIDsStmt, listTaskParticipantsByTenTaskIDs,
		arg.WorkspaceID,
		arg.TaskID,
		arg.TaskID_2,
		arg.TaskID_3,
		arg.TaskID_4,
		arg.TaskID_5,
		arg.TaskID_6,
		arg.TaskID_7,
		arg.TaskID_8,
		arg.TaskID_9,
		arg.TaskID_10,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TaskParticipant
	for rows.Next() {
		var i TaskParticipant
		if err := rows.Scan(
			&i.WorkspaceID,
			&i.TaskID,
			&i.ParticipantID,
			&i.RoleType,
			&i.ParticipantType,
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

const listTaskParticipantsByThreeTaskIDs = `-- name: ListTaskParticipantsByThreeTaskIDs :many
SELECT workspace_id, task_id, participant_id, role_type, participant_type, created_at, updated_at
FROM task_participants
WHERE  workspace_id = ? AND task_id IN (?,?,?)
`

type ListTaskParticipantsByThreeTaskIDsParams struct {
	WorkspaceID string
	TaskID      int64
	TaskID_2    int64
	TaskID_3    int64
}

func (q *Queries) ListTaskParticipantsByThreeTaskIDs(ctx context.Context, arg ListTaskParticipantsByThreeTaskIDsParams) ([]TaskParticipant, error) {
	rows, err := q.query(ctx, q.listTaskParticipantsByThreeTaskIDsStmt, listTaskParticipantsByThreeTaskIDs,
		arg.WorkspaceID,
		arg.TaskID,
		arg.TaskID_2,
		arg.TaskID_3,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TaskParticipant
	for rows.Next() {
		var i TaskParticipant
		if err := rows.Scan(
			&i.WorkspaceID,
			&i.TaskID,
			&i.ParticipantID,
			&i.RoleType,
			&i.ParticipantType,
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

const listTaskParticipantsByTwoTaskIDs = `-- name: ListTaskParticipantsByTwoTaskIDs :many
SELECT workspace_id, task_id, participant_id, role_type, participant_type, created_at, updated_at
FROM task_participants
WHERE  workspace_id = ? AND task_id IN (?,?)
`

type ListTaskParticipantsByTwoTaskIDsParams struct {
	WorkspaceID string
	TaskID      int64
	TaskID_2    int64
}

func (q *Queries) ListTaskParticipantsByTwoTaskIDs(ctx context.Context, arg ListTaskParticipantsByTwoTaskIDsParams) ([]TaskParticipant, error) {
	rows, err := q.query(ctx, q.listTaskParticipantsByTwoTaskIDsStmt, listTaskParticipantsByTwoTaskIDs, arg.WorkspaceID, arg.TaskID, arg.TaskID_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TaskParticipant
	for rows.Next() {
		var i TaskParticipant
		if err := rows.Scan(
			&i.WorkspaceID,
			&i.TaskID,
			&i.ParticipantID,
			&i.RoleType,
			&i.ParticipantType,
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
