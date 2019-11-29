package main

import (
	_ "github.com/kanhaiya15/gopf/conf"
	"github.com/kanhaiya15/gopf/conf/dbs/gopfmysql"
	"github.com/kanhaiya15/gopf/conf/dbs/gopfredis"

	"github.com/kanhaiya15/gopf/server/gopfserver"
)

func main() {
	mysqlChan := make(chan bool)
	redisChan := make(chan bool)
	go gopfmysql.Init(mysqlChan)
	go gopfredis.Init(redisChan)
	<-mysqlChan
	<-redisChan
	gopfserver.Init()
}
