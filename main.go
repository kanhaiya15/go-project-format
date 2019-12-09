package main

import (
	_ "github.com/kanhaiya15/gopf/conf"
	"github.com/kanhaiya15/gopf/conf/dbs/gopfmysql"
	"github.com/kanhaiya15/gopf/conf/dbs/gopfredis"
	"github.com/kanhaiya15/gopf/constants"
	"github.com/kanhaiya15/gopf/lib/consumer/kafkaconsumer"
	"github.com/kanhaiya15/gopf/lib/consumer/redisconsumer"

	"github.com/kanhaiya15/gopf/server/gopfserver"
)

func main() {
	mysqlChan := make(chan struct{})
	redisChan := make(chan struct{})
	go gopfmysql.Init(mysqlChan)
	go gopfredis.Init(redisChan)
	<-mysqlChan
	<-redisChan
	conn, _ := gopfredis.GetConn()
	go redisconsumer.InitConsumer(conn)
	go kafkaconsumer.InitConsumer(constants.KafkaBroker)
	gopfserver.Init()
}
