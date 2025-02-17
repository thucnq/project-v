// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: read_ticket_flowchart_history.sql

package sqlc

import (
	"context"
)

const listTicketFlowchartHistories = `-- name: ListTicketFlowchartHistories :many
SELECT workspace_id, ticket_id, id, flowchart, created_at
FROM ticket_flowchart_histories
WHERE workspace_id = ? AND ticket_id = ?
`

type ListTicketFlowchartHistoriesParams struct {
	WorkspaceID string
	TicketID    int64
}

func (q *Queries) ListTicketFlowchartHistories(ctx context.Context, arg ListTicketFlowchartHistoriesParams) ([]TicketFlowchartHistory, error) {
	rows, err := q.query(ctx, q.listTicketFlowchartHistoriesStmt, listTicketFlowchartHistories, arg.WorkspaceID, arg.TicketID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TicketFlowchartHistory
	for rows.Next() {
		var i TicketFlowchartHistory
		if err := rows.Scan(
			&i.WorkspaceID,
			&i.TicketID,
			&i.ID,
			&i.Flowchart,
			&i.CreatedAt,
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
