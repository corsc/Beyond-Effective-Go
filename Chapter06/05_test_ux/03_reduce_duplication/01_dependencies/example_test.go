package _1_dependencies

import (
	"context"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func TestUsingRedis(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	redisConn := getRedisConn(t, ctx, ":6379")

	// use Redis in test code

	// fake code to make the compiler happy
	assert.NotNil(t, redisConn)
}

func getRedisConn(t *testing.T, ctx context.Context, host string) redis.Conn {
	client, err := redis.DialContext(ctx, "tcp", host)
	if err != nil {
		t.Skipf("skipped due to failure to init Redis. err: %s", err)
	}

	return client
}
