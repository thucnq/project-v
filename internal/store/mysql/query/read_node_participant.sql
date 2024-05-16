-- name: ListNodeParticipantsByNodeID :many
SELECT *
FROM node_participants
WHERE workspace_id = ? AND node_id = ?;

-- name: ListPreviewNodeAssignees :many
SELECT ranked.participant_id, ranked.participant_type, ranked.workspace_id, ranked.node_id, participants.name AS participant_name
FROM (
         SELECT *, @row_number := IF(@node_id = node_id, @row_number + 1, 1) AS row_num, @node_id := node_id AS dummy
         FROM node_participants
         WHERE node_participants.workspace_id = ? AND node_id IN (
             SELECT nodes.id
             FROM nodes
             WHERE nodes.workspace_id = ? AND workflow_id = ?
             ) AND role_type = 4
         ORDER BY node_id, participant_id
     ) AS ranked
LEFT JOIN participants ON ranked.participant_id = participants.id
WHERE row_num <= 5;

-- name: ListPreviewNodeAdmins :many
SELECT ranked.participant_id, ranked.participant_type, ranked.workspace_id, ranked.node_id, participants.name AS participant_name
FROM (
         SELECT *, @row_number := IF(@node_id = node_id, @row_number + 1, 1) AS row_num, @node_id := node_id AS dummy
         FROM node_participants
         WHERE node_participants.workspace_id = ? AND node_id IN (
             SELECT nodes.id
             FROM nodes
             WHERE nodes.workspace_id = ? AND workflow_id = ?
             ) AND role_type = 1
         ORDER BY node_id, participant_id
     ) AS ranked
         LEFT JOIN participants ON ranked.participant_id = participants.id
WHERE row_num <= 5;
