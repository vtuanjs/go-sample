package impl

import (
	"context"
	"vtuanjs/my-project/internal/database"
)

type eUserInfo struct {
	r *database.Queries
}

func NewUserInfo(r *database.Queries) *eUserInfo {
	return &eUserInfo{r: r}
}

func (s *eUserInfo) GetInfoById(ctx context.Context, id string) error {
	return nil
}

func (s *eUserInfo) GetAllUser(ctx context.Context) error {
	return nil
}
