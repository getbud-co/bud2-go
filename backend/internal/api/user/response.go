package user

import (
	"time"

	appuser "github.com/getbud-co/bud2/backend/internal/app/user"
)

type Response struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Role          string `json:"role"`
	Status        string `json:"status"`
	IsSystemAdmin bool   `json:"is_system_admin"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type ListResponse struct {
	Data  []Response `json:"data"`
	Total int64      `json:"total"`
	Page  int        `json:"page"`
	Size  int        `json:"size"`
}

func toResponse(m *appuser.Member) Response {
	return Response{
		ID:            m.User.ID.String(),
		Name:          m.User.Name,
		Email:         m.User.Email,
		Role:          string(m.MembershipRole),
		Status:        string(m.MembershipStatus),
		IsSystemAdmin: m.User.IsSystemAdmin,
		CreatedAt:     m.User.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     m.User.UpdatedAt.Format(time.RFC3339),
	}
}
