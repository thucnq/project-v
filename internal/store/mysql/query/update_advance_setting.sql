-- name: UpdateAdvanceSetting :execresult
UPDATE advance_settings SET
	is_published=?,
	allow_cancel_value=?,
	allow_cancel_participant_value=?,
	allow_delete_value=?,
	allow_delete_participant_value=?,
	close_option_value=?,
	close_option_participant_value=?,
	approve_before_on_hold_value=?,
	approve_before_on_hold_participant_value=?,
	rate_option_value=?,
	version=?,
	updated_by=?,
	updated_at=?
WHERE workspace_id=? AND workflow_id=?;
