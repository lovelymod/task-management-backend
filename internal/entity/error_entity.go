package entity

import "errors"

var (
	// Global Errors
	ErrGlobalServerError            = errors.New("internal_server_error")
	ErrGlobalNotFound               = errors.New("not_found")
	ErrGlobalFileSizeExceedLimit    = errors.New("file_size_exceed_limit")
	ErrGlobalInvalidFileContentType = errors.New("invalid_file_content_type")
	ErrGlobalNotHavePermission      = errors.New("not_have_permission")

	// Auth Errors
	ErrAuthThisEmailIsAlreadyUsed = errors.New("this_email_is_already_used")
	ErrAuthWrongEmailOrPassword   = errors.New("wrong_email_or_password")

	ErrAuthAccessTokenExpired     = errors.New("access_token_expired")
	ErrAuthAccessTokenInvalid     = errors.New("access_token_invalid")
	ErrAuthAccessTokenNotProvided = errors.New("access_token_not_provided")

	ErrAuthRefreshTokenExpired     = errors.New("refresh_token_expired")
	ErrAuthRefreshTokenInvalid     = errors.New("refresh_token_invalid")
	ErrAuthRefreshTokenNotProvided = errors.New("refresh_token_not_provided")

	// Project Errors
	ErrProjectInvalidProjectId    = errors.New("invalid_project_id")
	ErrProjectInvalidStatusId     = errors.New("invalid_status_id")
	ErrProjectProjectIdIsRequired = errors.New("project_id_is_required")
	ErrProjectStatusIdIsRequired  = errors.New("status_id_is_required")
	ErrProjectDuplicateStatusName = errors.New("duplicate_status_name")
)
