package main

import (
	"SunCache/cache/log"
	"SunCache/cache/member"
	"SunCache/cache/team"
	"SunCache/data"
	"SunCache/data/file"
	"flag"
	"fmt"
)

var Db data.Data

func init() {
	Db = file.NewFileDb(file.ModeStream)
}

func main() {
	var port int
	var api bool
	var apiPort int
	flag.IntVar(&port, "port", 8333, "SunCache server port")
	flag.BoolVar(&api, "api", false, "Start api server?")
	flag.IntVar(&apiPort, "api-port", 8300, "Api server port")
	flag.Parse()

	socket := fmt.Sprintf("localhost:%v", port)
	addresses := []string{
		"localhost:8333",
		"localhost:8334",
		"localhost:8335",
	}
	team := team.NewTeam(
		socket,
		addresses...,
	)

	member := member.NewMember(
		"userInfo",
		2<<10,
		member.SourceGetterFunc(func(key string) ([]byte, error) {
			log.Info("[%v] DB search key: %v\n", socket, key)
			v, err := Db.Get(key)
			if err != nil {
				return nil, err
			} else if v == nil {
				return nil, fmt.Errorf("[%s] %s not found", socket, key)
			}
			return v, nil
		}),
		nil,
	)

	team.AddMember(member)
	if api {
		apiSocket := fmt.Sprintf("localhost:%v", apiPort)
		go team.StartHttpServer(apiSocket)
	}

	team.RunChatServer()
}
