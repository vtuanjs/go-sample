package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"vtuanjs/my-project/global"

	"github.com/redis/go-redis/v9"
)

func GetCache(ctx context.Context, key string, obj interface{}) error {
	rs, err := global.RDB.Get(ctx, key).Result()
	if err == redis.Nil {
		return fmt.Errorf("key %s not found", key)
	}

	// convert json to object
	// []byte(rs) = Buffer.from(rs) in TypeScript
	if err := json.Unmarshal([]byte(rs), obj); err != nil {
		return fmt.Errorf("failed to unmarshal json: %s", err)
	}

	return nil
}
