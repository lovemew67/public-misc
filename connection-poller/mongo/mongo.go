package mongo

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/lovemew67/cornerstone"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	connectionPoller = "connection-poller"
)

var (
	client    *mongo.Client
	ctx       = cornerstone.NewContext()
	nativeCtx = &ctx
)

func SimpleInit(username, password, authDatabase string, hosts []string) (err error) {
	account := ""
	if username != "" && password != "" {
		account = fmt.Sprintf("%s:%s@", url.QueryEscape(username), url.QueryEscape(password))
	}
	uri := fmt.Sprintf("mongodb://%s%s/%s", account, strings.Join(hosts, ","), authDatabase)
	clientOpt := options.Client().ApplyURI(uri)
	clientOpt.SetAppName(connectionPoller)
	clientOpt.SetSocketTimeout(5 * time.Second)
	clientOpt.SetMaxPoolSize(uint64(10))
	clientOpt.SetMinPoolSize(uint64(10))

	readPref, err := readpref.New(readpref.PrimaryMode)
	if err != nil {
		return
	}

	clientOpt.SetReadPreference(readPref)

	client, err = mongo.Connect(nativeCtx, clientOpt)
	if err != nil {
		return
	}

	err = client.Ping(nativeCtx, nil)
	if err != nil {
		_ = client.Disconnect(nativeCtx)
	}
	return
}

func CountDocuments(database, collection string, selector bson.M) (result int, err error) {
	col := client.Database(database).Collection(collection)
	count, err := col.CountDocuments(nativeCtx, selector)
	result = int(count)
	return
}
