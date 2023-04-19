package fruits

import (
	"context"
	"encoding/base64"
	"fmt"
	"fruits_microservice/config"

	"github.com/redis/go-redis/v9"
)

// Updates the fruit for a given user
// There's no validation here, the record is just updated
// straight away with what you put into this function!
// Error should be treated as >=500
func UpdateFruit(key string, username string, fruit string) error {
	client := redis.NewClient(&redis.Options{
		Addr: config.RedisEndpoint,
	})
	ctx := context.TODO()
	if client == nil || client.Ping(ctx).Err() != nil {
		return fmt.Errorf("cannot connect to redis, aborting")
	}

	username_b64 := base64.StdEncoding.EncodeToString([]byte(username))
	value := fmt.Sprintf("%s:%s", username_b64, fruit)

	return client.Set(ctx, key, value, 0).Err()
}
