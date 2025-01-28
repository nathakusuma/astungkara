package errorpkg

import (
	"net/http"
)

var (
	ErrInternalServer = NewError(http.StatusInternalServerError).
				WithErrorCode("INTERNAL_SERVER_ERROR").
				WithMessage("Something went wrong in our server. Please try again later.")

	ErrConferenceEnded = NewError(http.StatusUnprocessableEntity).
				WithErrorCode("CONFERENCE_ENDED").
				WithMessage("Conference has ended. You're not allowed to register anymore.")

	ErrConferenceFull = NewError(http.StatusConflict).
				WithErrorCode("CONFERENCE_FULL").
				WithMessage("Conference is full. You're not allowed to register anymore.")

	ErrConflictingRegistrations = NewError(http.StatusConflict).
					WithErrorCode("CONFLICTING_REGISTRATIONS").
					WithMessage("You have conflicting registrations. Please check your schedule.")

	ErrCredentialsNotMatch = NewError(http.StatusUnauthorized).
				WithErrorCode("CREDENTIALS_NOT_MATCH").
				WithMessage("Credentials do not match. Please try again.")

	ErrEmailAlreadyRegistered = NewError(http.StatusConflict).
					WithErrorCode("EMAIL_ALREADY_REGISTERED").
					WithMessage("Email already registered. Please login or use another email.")

	ErrEndTimeBeforeStart = NewError(http.StatusUnprocessableEntity).
				WithErrorCode("END_TIME_BEFORE_START").
				WithMessage("End time is before start time. Please use correct time.")

	ErrFailParseRequest = NewError(http.StatusBadRequest).
				WithErrorCode("FAIL_PARSE_REQUEST").
				WithMessage("Failed to parse request. Please check your request format.")

	ErrForbiddenRole = NewError(http.StatusForbidden).
				WithErrorCode("FORBIDDEN_ROLE").
				WithMessage("You're not allowed to access this resource.")

	ErrForbiddenUser = NewError(http.StatusForbidden).
				WithErrorCode("FORBIDDEN_USER").
				WithMessage("You're not allowed to access this resource.")

	ErrHostCannotRegister = NewError(http.StatusForbidden).
				WithErrorCode("HOST_CANNOT_REGISTER").
				WithMessage("You're not allowed to register to your own conference.")

	ErrInvalidBearerToken = NewError(http.StatusUnauthorized).
				WithErrorCode("INVALID_BEARER_TOKEN").
				WithMessage("Your auth session is invalid. Please renew your auth session.")

	ErrInvalidOTP = NewError(http.StatusUnauthorized).
			WithErrorCode("INVALID_OTP").
			WithMessage("Invalid OTP. Please try again or request a new OTP.")

	ErrInvalidPagination = NewError(http.StatusUnprocessableEntity).
				WithErrorCode("INVALID_PAGINATION").
				WithMessage("Cannot use after_id and before_id at the same time.")

	ErrInvalidRefreshToken = NewError(http.StatusUnauthorized).
				WithErrorCode("INVALID_REFRESH_TOKEN").
				WithMessage("Auth session is invalid. Please login again.")

	ErrNoBearerToken = NewError(http.StatusUnauthorized).
				WithErrorCode("NO_BEARER_TOKEN").
				WithMessage("You're not logged in. Please login first.")

	ErrNotFound = NewError(http.StatusNotFound).
			WithErrorCode("NOT_FOUND").
			WithMessage("Data not found.")

	ErrTimeAlreadyPassed = NewError(http.StatusUnprocessableEntity).
				WithErrorCode("TIME_ALREADY_PASSED").
				WithMessage("Time has already passed. Please use future time.")

	ErrTimeWindowConflict = NewError(http.StatusConflict).
				WithErrorCode("TIME_WINDOW_CONFLICT").
				WithMessage("There's already a conference in the same time window. Please choose another time window.")

	ErrUpdateApprovedTimeWindow = NewError(http.StatusUnprocessableEntity).
					WithErrorCode("UPDATE_APPROVED_TIME_WINDOW").
					WithMessage("You're not allowed to update an approved conference time window.")

	ErrUpdatePastConference = NewError(http.StatusUnprocessableEntity).
				WithErrorCode("UPDATE_PAST_CONFERENCE").
				WithMessage("You're not allowed to update a past conference.")

	ErrUpdateRejectedConference = NewError(http.StatusUnprocessableEntity).
					WithErrorCode("UPDATE_REJECTED_CONFERENCE").
					WithMessage("You're not allowed to update a rejected conference.")

	ErrUpdatePastConferenceStatus = NewError(http.StatusUnprocessableEntity).
					WithErrorCode("UPDATE_PAST_CONFERENCE_STATUS").
					WithMessage("You're only allowed to change past conference status from pending to rejected.")

	ErrUserHasActiveProposal = NewError(http.StatusConflict).
					WithErrorCode("USER_HAS_ACTIVE_PROPOSAL").
					WithMessage("You already have an active proposal. Please wait until it's accepted or delete it.")

	ErrUserAlreadyRegisteredToConference = NewError(http.StatusConflict).
						WithErrorCode("USER_ALREADY_REGISTERED_TO_CONFERENCE").
						WithMessage("You're already registered to this conference.")
)
