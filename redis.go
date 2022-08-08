package foxkit

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v9"
)

// connects to the redis server specified in ENV, tests the connection
func ConnectRedis(ctx context.Context, dbNumber int) *redis.Client {
	client := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_HOST"), Password: os.Getenv("REDIS_PASS"), DB: dbNumber})
	// ping the server
	lctx, cancel := context.WithTimeout(ctx, time.Second*10)
	if err := client.Ping(lctx).Err(); err != nil {
		ErrorFatal("FoxKit", err)
	}
	cancel()
	// connection is working, return the client
	return client
}
