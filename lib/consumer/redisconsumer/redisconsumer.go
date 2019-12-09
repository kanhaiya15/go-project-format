package redisconsumer

import (
	"github.com/gomodule/redigo/redis"
	"github.com/kanhaiya15/gopf/lib/logging/gopflogrus"
)

var logger = gopflogrus.NewLogger()

// InitConsumer InitConsumer
func InitConsumer(rc redis.Conn) {
	psc := redis.PubSubConn{Conn: rc}
	psc.PSubscribe("__keyevent@0__:expired")
	for {
		switch msg := psc.Receive().(type) {
		case redis.Message:
			logger.Infof("Redis Consumer Message: %+v", msg.Data)
		case redis.Subscription:
			logger.Infof("Redis Consumer Subscription : %s %s %d\n", msg.Kind, msg.Channel, msg.Count)
		case error:
			logger.Infof("Redis Consumer error : %v\n", msg)
		}
	}
}
