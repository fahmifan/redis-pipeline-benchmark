# Sucess Matrix

- Improve performance by reduce latency

## How To

- simple bench: `make bench`
- detail bench: `make bench-full`
  - this benchmarks based on [go_cpu_memory_profiling_benchmarks.sh](https://gist.github.com/arsham/bbc93990d8e5c9b54128a3d88901ab90)

## Benchmarks
### TLDR
Performance increase in all benchmarks when using `redigo` pipeline.

Tested on
```
MacBook Pro (13-inch, 2018)
2,3 GHz Quad-Core Intel Core i5
8 GB 2133 MHz LPDDR3
```

## Result

```
# make bench

goos: darwin
goarch: amd64
pkg: github.com/fahmifan/redis-pipeline-benchmark
Benchmark_LocalRedis/pipes__batch_10-8         	    8773	    116976 ns/op
Benchmark_LocalRedis/nromal_batch_10-8         	    1332	   1150114 ns/op
Benchmark_LocalRedis/pipes__batch_100-8        	    3388	    360788 ns/op
Benchmark_LocalRedis/nromal_batch_100-8        	     135	   9446956 ns/op
Benchmark_LocalRedis/pipes__batch_1000-8       	     662	   1876936 ns/op
Benchmark_LocalRedis/nromal_batch_1000-8       	      13	  89030635 ns/op
Benchmark__MiniRedis/pipes__batch_10-8         	   10000	    121230 ns/op
Benchmark__MiniRedis/nromal_batch_10-8         	    2044	    602280 ns/op
Benchmark__MiniRedis/pipes__batch_100-8        	    1404	    830698 ns/op
Benchmark__MiniRedis/nromal_batch_100-8        	     200	   6156052 ns/op
Benchmark__MiniRedis/pipes__batch_1000-8       	     142	   8044406 ns/op
Benchmark__MiniRedis/nromal_batch_1000-8       	      19	  59559810 ns/op
PASS
ok  	github.com/fahmifan/redis-pipeline-benchmark	18.872s
```

```
# benchstat bench.txt

name                             time/op
_LocalRedis/pipes__batch_10-8     105µs ± 7%
_LocalRedis/nromal_batch_10-8     894µs ± 4%
_LocalRedis/pipes__batch_100-8    231µs ± 0%
_LocalRedis/nromal_batch_100-8   10.2ms ±19%
_LocalRedis/pipes__batch_1000-8  1.24ms ± 2%
_LocalRedis/nromal_batch_1000-8  96.0ms ± 3%
__MiniRedis/pipes__batch_10-8     167µs ± 5%
__MiniRedis/nromal_batch_10-8     940µs ± 2%
__MiniRedis/pipes__batch_100-8    947µs ± 4%
__MiniRedis/nromal_batch_100-8   9.17ms ± 7%
__MiniRedis/pipes__batch_1000-8  9.59ms ± 3%
__MiniRedis/nromal_batch_1000-8  87.3ms ± 4%

name                             alloc/op
_LocalRedis/pipes__batch_10-8    1.17kB ± 0%
_LocalRedis/nromal_batch_10-8    1.60kB ± 0%
_LocalRedis/pipes__batch_100-8   12.7kB ± 0%
_LocalRedis/nromal_batch_100-8   16.0kB ± 0%
_LocalRedis/pipes__batch_1000-8   128kB ± 0%
_LocalRedis/nromal_batch_1000-8   168kB ± 0%
__MiniRedis/pipes__batch_10-8    2.96kB ± 0%
__MiniRedis/nromal_batch_10-8    4.48kB ± 0%
__MiniRedis/pipes__batch_100-8   30.3kB ± 0%
__MiniRedis/nromal_batch_100-8   44.8kB ± 0%
__MiniRedis/pipes__batch_1000-8   304kB ± 0%
__MiniRedis/nromal_batch_1000-8   464kB ± 0%

name                             allocs/op
_LocalRedis/pipes__batch_10-8      61.0 ± 0%
_LocalRedis/nromal_batch_10-8      70.0 ± 0%
_LocalRedis/pipes__batch_100-8      601 ± 0%
_LocalRedis/nromal_batch_100-8      700 ± 0%
_LocalRedis/pipes__batch_1000-8   6.00k ± 0%
_LocalRedis/nromal_batch_1000-8   7.00k ± 0%
__MiniRedis/pipes__batch_10-8       231 ± 0%
__MiniRedis/nromal_batch_10-8       330 ± 0%
__MiniRedis/pipes__batch_100-8    2.21k ± 0%
__MiniRedis/nromal_batch_100-8    3.30k ± 0%
__MiniRedis/pipes__batch_1000-8   22.0k ± 0%
__MiniRedis/nromal_batch_1000-8   33.0k ± 0%
```