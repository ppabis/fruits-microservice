package integration

import "math/rand"

func Serve() (int, error) {
	port := 58000 + rand.Intn(1000)

	return port, nil
}
