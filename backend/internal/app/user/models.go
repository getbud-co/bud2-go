package user

import (
	"strings"

	"github.com/google/uuid"

	"github.com/getbud-co/bud2/backend/internal/domain/membership"
	usr "github.com/getbud-co/bud2/backend/internal/domain/user"
)

type Member struct {
	User             usr.User
	OrganizationID   uuid.UUID
	MembershipRole   membership.Role
	MembershipStatus membership.Status
}

type MemberListResult struct {
	Members []Member
	Total   int64
}

func emailDomain(email string) string {
	parts := strings.Split(strings.TrimSpace(email), "@")
	if len(parts) != 2 {
		return ""
	}
	return strings.ToLower(parts[1])
}
