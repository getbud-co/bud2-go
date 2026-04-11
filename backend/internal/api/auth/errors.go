package auth

import (
	"errors"
	"net/http"

	"github.com/dsbraz/bud2/backend/internal/api/httputil"
	appauth "github.com/dsbraz/bud2/backend/internal/app/auth"
)

func handleAuthError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, appauth.ErrInvalidCredentials):
		httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "invalid credentials")
	case errors.Is(err, appauth.ErrUserInactive):
		httputil.WriteProblem(w, http.StatusForbidden, "Forbidden", "user account is inactive")
	case errors.Is(err, appauth.ErrNoOrganizations):
		httputil.WriteProblem(w, http.StatusForbidden, "Forbidden", "user has no accessible organizations")
	default:
		httputil.WriteProblem(w, http.StatusInternalServerError, "Internal Server Error", "an unexpected error occurred")
	}
}
