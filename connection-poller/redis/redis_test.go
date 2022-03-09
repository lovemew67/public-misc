package redis

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
)

var (
	dockerPool     *dockertest.Pool
	dockerResource *dockertest.Resource
)

func beforeTest() {
	// init error
	var err error

	// init docker
	dockerPool, err = dockertest.NewPool("")
	if err != nil {
		panic("docker test init fail, error:" + err.Error())
	}
	runOpts := &dockertest.RunOptions{
		Name:       "",
		Repository: "redis",
		Tag:        "5.0",
		Env:        []string{},
	}
	hcOpts := func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	}
	dockerResource, err = dockerPool.RunWithOptions(runOpts, hcOpts)
	if err != nil {
		panic("redis docker init fail, error:" + err.Error())
	}
	err = dockerResource.Expire(600)
	if err != nil {
		panic("failed to set expire, err: " + err.Error())
	}
	err = dockerPool.Retry(func() error {
		// test connection or setup
		var errNew error
		_, errNew = NewPool(&Config{
			Host:    fmt.Sprintf("localhost:%s", dockerResource.GetPort("6379/tcp")),
			RedisDB: 1,
		})
		return errNew
	})
	if err != nil {
		panic("connect docker fail, error:" + err.Error())
	}
}

func TestMain(m *testing.M) {
	log.SetFlags(log.LstdFlags)
	log.SetOutput(os.Stderr)
	beforeTest()
	retCode := m.Run()
	afterTest()
	os.Exit(retCode)
}

func afterTest() {
	_ = dockerPool.Purge(dockerResource)
}

func Test_All(t *testing.T) {
	// test: init incorrect pool
	cfg := &Config{
		Host:    "localhost:9453",
		RedisDB: 1,
	}
	pool, err := NewPool(cfg)
	assert.Nil(t, pool)
	assert.Error(t, err)

	// test: init correct pool
	cfg = &Config{
		Host:    fmt.Sprintf("localhost:%s", dockerResource.GetPort("6379/tcp")),
		RedisDB: 1,
	}
	pool, err = NewPool(cfg)
	assert.NotNil(t, pool)
	assert.NoError(t, err)

	// test: exist before set
	testKey := "test-key"
	exist, err := pool.Exists(testKey)
	assert.NoError(t, err)
	assert.False(t, exist)

	// test: get before set

	value, err := pool.GetBytes(testKey)
	assert.Error(t, err)
	assert.Equal(t, "redigo: nil returned", err.Error())
	assert.Equal(t, 0, len(value))

	// test: delete before set
	err = pool.Delete(testKey)
	assert.Error(t, err)

	// test: set
	testValue := "test-value"
	err = pool.Set(testKey, testValue)
	assert.NoError(t, err)

	// test: exist after set
	exist, err = pool.Exists(testKey)
	assert.NoError(t, err)
	assert.True(t, exist)

	// test: get after set
	value, err = pool.GetBytes(testKey)
	assert.NoError(t, err)
	assert.Equal(t, testValue, string(value))

	// test: delete after set
	err = pool.Delete(testKey)
	assert.NoError(t, err)

	// test: exist after delete
	exist, err = pool.Exists(testKey)
	assert.NoError(t, err)
	assert.False(t, exist)

	// test: get after delete
	value, err = pool.GetBytes(testKey)
	assert.Error(t, err)
	assert.Equal(t, "redigo: nil returned", err.Error())
	assert.Equal(t, 0, len(value))
}
