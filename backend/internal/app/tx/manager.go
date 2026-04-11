// Package tx defines application-layer transaction contracts shared by use cases.
package tx

import (
	"context"

	"github.com/dsbraz/bud2/backend/internal/domain/membership"
	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/dsbraz/bud2/backend/internal/domain/user"
)

type Repositories interface {
	Organizations() organization.Repository
	Users() user.Repository
	Memberships() membership.Repository
}

type Manager interface {
	WithTx(ctx context.Context, fn func(repos Repositories) error) error
}
