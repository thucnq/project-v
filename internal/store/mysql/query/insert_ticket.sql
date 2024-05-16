-- name: CreateTicket :execresult
INSERT INTO tickets (
    workspace_id, workflow_id, id, code, title, priority,
    is_private, status, ref_ticket_ids, workflow, rating_point,
    review, created_by, updated_by, created_at, updated_at
)
VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);