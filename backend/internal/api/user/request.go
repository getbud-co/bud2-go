package user

import (
	"github.com/google/uuid"

	"github.com/dsbraz/bud2/backend/internal/app/user"
	"github.com/dsbraz/bud2/backend/internal/domain"
)

type createRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"required,oneof=admin manager collaborator"`
}

type updateRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=100"`
	Email  string `json:"email" validate:"required,email"`
	Role   string `json:"role" validate:"required,oneof=admin manager collaborator"`
	Status string `json:"status" validate:"required,oneof=invited active inactive"`
}

func (r createRequest) toCommand(organizationID domain.TenantID) user.CreateCommand {
	return user.CreateCommand{
		OrganizationID: organizationID,
		Name:           r.Name,
		Email:          r.Email,
		Password:       r.Password,
		Role:           r.Role,
	}
}

func (r updateRequest) toCommand(organizationID domain.TenantID, id uuid.UUID) user.UpdateCommand {
	return user.UpdateCommand{
		OrganizationID: organizationID,
		ID:             id,
		Name:           r.Name,
		Email:          r.Email,
		Role:           r.Role,
		Status:         r.Status,
	}
}
