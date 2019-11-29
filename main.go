package main

import (
	_ "github.com/kanhaiya15/gopf/conf"
	"github.com/kanhaiya15/gopf/conf/dbs/gopfmysql"
	"github.com/kanhaiya15/gopf/conf/dbs/gopfredis"

	"github.com/kanhaiya15/gopf/server/gopfserver"
)

func main() {
	gopfmysql.Init()
	gopfredis.Init()
	gopfserver.Init()
}
