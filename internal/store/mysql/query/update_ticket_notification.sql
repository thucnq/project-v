-- name: UpdateTicketNotification :execresult
UPDATE ticket_notifications SET
	is_published=?,

	created_request_value=?,
	created_request_participant_value=?,

	updated_asignee_value=?,
	updated_asignee_participant_value=?,

	updated_ticket_value=?,
	updated_ticket_participant_value=?,

	updated_workflow_value=?,
	updated_workflow_participant_value=?,

	updated_status_step_value=?,
	updated_status_step_participant_value=?,

	updated_comment_value=?,
	updated_comment_participant_value=?,

	updated_follower_or_supporter_value=?,
	updated_follower_or_supporter_participant_value=?,

	updated_rating_service_value=?,
	updated_rating_service_participant_value=?,

	almost_expired_first_response_value=?,
	almost_expired_first_response_participant_value=?,

	almost_expired_process_ticket_value=?,
	almost_expired_process_ticket_participant_value=?,

	expired_first_response_value=?,
	expired_first_response_participant_value=?,

	expired_process_ticket_value=?,
	expired_process_ticket_participant_value=?,

	version=?,
	updated_by=?,
	updated_at=?
WHERE workspace_id=? AND id=?;
