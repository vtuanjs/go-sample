package impl

import (
	"context"
	"vtuanjs/my-project/internal/database"
)

type eUserAdmin struct {
	r *database.Queries
}

func NewUserAdmin(r *database.Queries) *eUserAdmin {
	return &eUserAdmin{r: r}
}

func (s *eUserAdmin) RemoveUser(ctx context.Context, id string) error {
	return nil
}

func (s *eUserAdmin) FindOneUser(ctx context.Context, id string) error {
	return nil
}
