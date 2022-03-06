package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		Repository: "mongo",
		Tag:        "4.2",
		Env: []string{
			"MONGO_INITDB_ROOT_USERNAME=username",
			"MONGO_INITDB_ROOT_PASSWORD=password",
		},
	}
	hcOpts := func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	}
	dockerResource, err = dockerPool.RunWithOptions(runOpts, hcOpts)
	if err != nil {
		panic("mongodb docker init fail, error:" + err.Error())
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

	// init client
	hosts := []string{
		fmt.Sprintf("localhost:%s", dockerResource.GetPort("27017/tcp")),
	}
	account := fmt.Sprintf("%s:%s@", "username", "password")
	uri := fmt.Sprintf("mongodb://%s%s/%s", account, strings.Join(hosts, ","), "admin")
	clientOpt := options.Client().ApplyURI(uri)
	client, err = mongo.Connect(nativeCtx, clientOpt)
	if err != nil {
		panic("connect mongo fail, error:" + err.Error())
	}

	// setup data
	_, err = client.Database("database").Collection("collection").InsertOne(context.Background(), bson.M{
		"id": "aaa",
	})
	if err != nil {
		panic("prepare data fail, error:" + err.Error())
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
	err := SimpleInit("username", "password", "admin", []string{
		fmt.Sprintf("localhost:%s", dockerResource.GetPort("27017/tcp")),
	})
	assert.NoError(t, err)

	result, err := CountDocuments("database", "collection", bson.M{})
	assert.NoError(t, err)
	assert.Equal(t, result, 1)
}
