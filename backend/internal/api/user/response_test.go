package user

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	appuser "github.com/dsbraz/bud2/backend/internal/app/user"
	"github.com/dsbraz/bud2/backend/internal/domain/membership"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
)

func TestToResponse(t *testing.T) {
	member := &appuser.Member{
		User:             *fixtures.NewUser(),
		MembershipRole:   membership.RoleAdmin,
		MembershipStatus: membership.StatusActive,
	}

	resp := toResponse(member)

	assert.Equal(t, member.User.ID.String(), resp.ID)
	assert.Equal(t, member.User.Name, resp.Name)
	assert.Equal(t, member.User.Email, resp.Email)
	assert.Equal(t, "admin", resp.Role)
	assert.Equal(t, "active", resp.Status)
	assert.Equal(t, member.User.IsSystemAdmin, resp.IsSystemAdmin)
	assert.Equal(t, member.User.CreatedAt.Format(time.RFC3339), resp.CreatedAt)
	assert.Equal(t, member.User.UpdatedAt.Format(time.RFC3339), resp.UpdatedAt)
}
