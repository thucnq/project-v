-- name: CreateAdvanceSetting :execresult
INSERT INTO advance_settings (
	workspace_id,
	workflow_id,
	is_published,
	created_by,
	updated_by,
	admin_option_participant_value,
	admin_option_participant_ref,
	allow_cancel_value,
	allow_cancel_participant_value,
	allow_cancel_participant_ref,
	allow_delete_value,
	allow_delete_participant_value,
	allow_delete_participant_ref,
	close_option_value,
	close_option_participant_value,
	close_option_participant_ref,
	approve_before_on_hold_value,
	approve_before_on_hold_participant_value,
	approve_before_on_hold_participant_ref,
	rate_option_value,
	version
)
VALUES (
	?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);
