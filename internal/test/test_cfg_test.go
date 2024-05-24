package test

import (
	"context"
	"testing"
	"time"

	"github.com/redis/rueidis"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestNewDynamicConfig(t *testing.T) {
	ctx := context.Background()

	r := require.New(t)

	redisContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "redis:7",
			ExposedPorts: []string{"6379/tcp"},
			WaitingFor:   wait.ForListeningPort("6379/tcp"),
		},
		Started: true,
	})
	require.NoError(t, err)

	defer func() {
		r.NoError(redisContainer.Terminate(ctx))
	}()

	redisHost, err := redisContainer.Host(ctx)
	r.NoError(err)

	redisPort, err := redisContainer.MappedPort(ctx, "6379")
	r.NoError(err)

	redisClient, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{redisHost + ":" + redisPort.Port()},
	})
	r.NoError(err)

	cfg := NewDynamicConfig(redisClient)
	testBool1 := cfg.TestBool1(ctx)
	r.True(testBool1)

	testBool2 := cfg.TestBool2(ctx)
	r.False(testBool2)

	err = cfg.StoreTestBool1(ctx, false)
	r.NoError(err)

	time.Sleep(time.Millisecond)

	testBool1 = cfg.TestBool1(ctx)
	r.False(testBool1)

	err = cfg.StoreTestBool1(ctx, true)
	r.NoError(err)

	time.Sleep(time.Millisecond)

	testBool1 = cfg.TestBool1(ctx)
	r.True(testBool1)
}
