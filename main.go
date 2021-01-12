package main

import (
	"fmt"
	"log"
	"time"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/kumparan/go-utils"
)

const redisAddr = "redis://localhost:6379"
const ndata = 100

func newRedisConnPool(url string) *redigo.Pool {
	return &redigo.Pool{
		MaxIdle:     100,
		MaxActive:   10000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.DialURL(url)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// User ..
type User struct {
	ID   int
	Name string
	Age  int
}

func userCacheKeyByID(id int) string {
	return fmt.Sprintf("user:%d", id)
}

func main() {
	redPool := newRedisConnPool("redis://localhost:6379/0")
	seed(redPool, ndata)
	withPipeline(redPool, ndata, true)
}

func withPipeline(redPool *redigo.Pool, ndata int, debug bool) {
	conn := redPool.Get()
	defer conn.Close()

	loop(ndata, func(val int) {
		conn.Send("GET", userCacheKeyByID(val))
	})
	conn.Flush()

	loop(ndata, func(val int) {
		repl, err := conn.Receive()
		if err != nil {
			if err == redigo.ErrNil {
				return
			}
			log.Fatal(err)
		}

		if debug {
			fmt.Println(string(repl.([]byte)))
		}
	})
}

func noPipeline(redPool *redigo.Pool, ndata int, debug bool) {
	loop(ndata, func(val int) {
		get(redPool, ndata, debug)
	})
}

func get(redPool *redigo.Pool, val int, debug bool) {
	conn := redPool.Get()
	defer conn.Close()

	rep, err := conn.Do("GET", userCacheKeyByID(val))
	if err != nil {
		if err == redigo.ErrNil {
			return
		}
		log.Fatal(err)
	}

	if debug {
		fmt.Println(string(rep.([]byte)))
	}
}

func seed(redPool *redigo.Pool, ndata int) {
	conn := redPool.Get()
	defer conn.Close()

	loop(ndata, func(i int) {
		user := &User{
			ID:   i,
			Name: fmt.Sprintf("John %d", i),
			Age:  23 % i,
		}

		_, err := conn.Do("SETEX", userCacheKeyByID(user.ID), 3600, utils.Dump(user))
		if err != nil {
			log.Fatal(err)
		}
	})
}

func loop(to int, fn func(val int)) {
	for i := 1; i <= to; i++ {
		fn(i)
	}
}
