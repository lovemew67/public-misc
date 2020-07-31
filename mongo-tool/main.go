package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/mongodb/mongo-tools-common/options"
	"github.com/mongodb/mongo-tools/mongoexport"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds)
}

func main() {
	opt := mongoexport.Options{
		&options.ToolOptions{
			// Auth: &options.Auth{
			// 	Username: "lcc5",
			// 	Password: "@lcc5",
			// 	Source:   "admin",
			// },
			// Connection: &options.Connection{
			// 	Host: "localhost",
			// 	Port: "27017",
			// },
			Namespace: &options.Namespace{
				DB:         "liveConnect",
				Collection: "tag",
			},
			// Verbosity: &options.Verbosity{},
			URI: &options.URI{
				ConnString: connstring.ConnString{
					AuthSource: "admin",
					Username:   "lcc5",
					Password:   "@lcc5",
					Hosts: []string{
						"localhost:27017",
					},
				},
			},
		},
		&mongoexport.OutputFormatOptions{
			Type:       "json",
			JSONFormat: "canonical",
		},
		&mongoexport.InputOptions{
			Query:   `{"service":{$ne:"cake-tw.m800.com"}}`,
			SlaveOk: false,
		},
		[]string{},
	}
	me, _ := mongoexport.New(opt)
	defer me.Close()
	out := &bytes.Buffer{}
	_, _ = me.Export(out)
	fmt.Println(out.String())
}
