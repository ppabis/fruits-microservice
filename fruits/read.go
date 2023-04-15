package fruits

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"fruits_microservice/config"
	"strings"

	"github.com/redis/go-redis/v9"
)

var ErrKeyNotFound = errors.New("key not found")

func GetFruits() (map[string]string, error) {
	client := redis.NewClient(&redis.Options{
		Addr: config.RedisEndpoint,
	})
	ctx := context.TODO()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("cannot connect to redis: %v", err)
	}

	fruits := make(map[string]string)
	iterator := client.Scan(ctx, 0, "user:*", 0).Iterator()
	for iterator.Next(ctx) {
		key := iterator.Val()
		record, err := client.Get(ctx, key).Result()

		if err != nil {
			return nil, fmt.Errorf("failed to get key %s: %v", key, err)
		}

		user, fruit, err := parseRecord(record)
		if err != nil {
			return nil, fmt.Errorf("failed to parse record: %v", err)
		}

		fruits[user] = fruit
	}

	if err := iterator.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate: %v", err)
	}

	return fruits, nil
}

func parseRecord(record string) (string, string, error) {
	// Returns username, fruit, nil or empty strings and error
	splitRecord := strings.Split(record, ":")
	if len(splitRecord) != 2 {
		return "", "", fmt.Errorf("invalid record, parts count is not 2")
	}

	decodedUsername, err := base64.StdEncoding.DecodeString(splitRecord[0])
	if err != nil {
		return "", "", fmt.Errorf("failed to decode username: %v", err)
	}

	return string(decodedUsername), splitRecord[1], nil
}

func GetFruit(id int) (string, string, error) {
	client := redis.NewClient(&redis.Options{
		Addr: config.RedisEndpoint,
	})
	ctx := context.TODO()

	if err := client.Ping(ctx).Err(); err != nil {
		return "", "", fmt.Errorf("cannot connect to redis: %v", err)
	}

	key := fmt.Sprintf("user:%d", id)
	record, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", "", ErrKeyNotFound
	}

	username, fruit, err := parseRecord(record)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse record: %v", err)
	}

	return username, fruit, nil
}
