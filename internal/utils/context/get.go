package utils

import (
	"context"
	"errors"
	"vtuanjs/my-project/internal/cache"
)

type InfoUserUUID struct {
	UserId      uint64
	UserAccount string
}

func GetSubjectUUID(ctx context.Context) (string, error) {
	sUUID, ok := ctx.Value("subjectUUID").(string)
	if !ok {
		return "", errors.New("subjectUUID not found")
	}
	return sUUID, nil
}

func GetUserFromUUID(ctx context.Context) (uint64, error) {
	sUUID, err := GetSubjectUUID(ctx)
	if err != nil {
		return 0, err
	}

	// get from cache
	var infoUser InfoUserUUID
	if err := cache.GetCache(ctx, sUUID, &infoUser); err != nil {
		return 0, err
	}
	return infoUser.UserId, nil
}
