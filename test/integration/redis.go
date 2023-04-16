package integration

import (
	"context"
	"math/rand"
	"strconv"

	dt "github.com/ory/dockertest/v3"
	d "github.com/ory/dockertest/v3/docker"
	"github.com/redis/go-redis/v9"
)

var pool *dt.Pool = nil

// Creates a temporary redis server with
// some test data
//
// user 1 - john, kiwi
// user 2 - johnathan, apple
// user 4 - damian, apple
// user 5 - alexis, pineapple
func RedisWithTestData() (*dt.Resource, int, error) {
	port := 55037 + rand.Intn(200)

	var err error
	pool, err = dt.NewPool("")
	if err != nil {
		return nil, -1, err
	}

	res, err := pool.RunWithOptions(
		&dt.RunOptions{
			Repository:   "redis",
			Tag:          "latest",
			ExposedPorts: []string{"6379/tcp"},
			PortBindings: map[d.Port][]d.PortBinding{
				"6379/tcp": {
					{
						HostPort: strconv.Itoa(port),
					},
				},
			},
		},
		func(config *d.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = d.RestartPolicy{Name: "no"}
		})

	if err != nil {
		return nil, -1, err
	}

	if err := pool.Retry(func() error {
		return redis.
			NewClient(&redis.Options{
				Addr: "localhost:" + strconv.Itoa(port),
			}).
			Ping(context.TODO()).
			Err()
	}); err != nil {
		return nil, -1, err
	}

	if err := testData(port); err != nil {
		return nil, -1, err
	}

	return res, port, nil
}

func DestroyContainer(r *dt.Resource) error {
	if r == nil {
		return nil
	}

	return pool.Purge(r)
}

func testData(port int) error {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:" + strconv.Itoa(port),
	})

	_, err := client.Pipelined(context.TODO(), func(pipe redis.Pipeliner) error {
		pipe.Set(context.TODO(), "user:1", "am9obg==:kiwi", 0)
		pipe.Set(context.TODO(), "user:2", "am9obmF0aGFu:apple", 0)
		pipe.Set(context.TODO(), "user:4", "ZGFtaWFu:apple", 0)
		pipe.Set(context.TODO(), "user:5", "YWxleGlz:pineapple", 0)
		return nil
	})

	return err
}
