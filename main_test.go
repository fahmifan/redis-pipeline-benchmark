package main

import (
	"fmt"
	"testing"

	"github.com/alicebob/miniredis/v2"
	redigo "github.com/gomodule/redigo/redis"
)

var _batches = []int{10, 100, 1000}

func Benchmark_LocalRedis(b *testing.B) {
	redPool := newRedisConnPool(redisAddr)
	for _, batch := range _batches {
		runBench(redPool, batch, b)
	}
}

func Benchmark__MiniRedis(b *testing.B) {
	mr, _ := miniredis.Run()
	defer mr.Close()
	redPool := newRedisConnPool("redis://" + mr.Addr())

	for _, batch := range _batches {
		runBench(redPool, batch, b)
	}
}

func runBench(redPool *redigo.Pool, batch int, b *testing.B) {
	seed(redPool, batch)

	b.Run(fmt.Sprintf("pipes__batch_%d", batch), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			withPipeline(redPool, batch, false)
		}
	})

	b.Run(fmt.Sprintf("nromal_batch_%d", batch), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			noPipeline(redPool, batch, false)
		}
	})
}
