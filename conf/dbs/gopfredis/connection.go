package gopfredis

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/kanhaiya15/gopf/utils"
)

// Pool
var (
	Pool *redis.Pool
)

// Init initialize
func Init() {
	Pool = newPool()
	err := Pool.TestOnBorrow(Pool.Get(), time.Time{})
	if err != nil {
		panic(err.Error())
	}
	cleanupHook()
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := getConnection()
			if err != nil {
				panic(err.Error())
			}
			c, err := redis.Dial("tcp", conn)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := ping(c)
			if err != nil {
				panic(err.Error())
			}
			return err
		},
	}
}

// ping tests connectivity for redis (PONG should be returned)
func ping(c redis.Conn) (string, error) {
	// Send PING command to Redis
	// PING command returns a Redis "Simple String"
	// Use redis.String to convert the interface type to string
	return redis.String(c.Do("PING"))
}

func getConnection() (string, error) {
	host, err := utils.GetConfValue("REDIS_HOST")
	if err != nil {
		panic(err.Error())
	}
	port, err := utils.GetConfValue("REDIS_PORT")
	if err != nil {
		panic(err.Error())
	}
	conn := fmt.Sprintf("%s:%s", host, port)
	return conn, nil
}

func cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}
