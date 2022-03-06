package redis

import (
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
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
		return nil
	})
	if err != nil {
		panic("connect docker fail, error:" + err.Error())
	}

	// dockerResource.GetPort("27017/tcp")
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

func Test_All(t *testing.T) {}
