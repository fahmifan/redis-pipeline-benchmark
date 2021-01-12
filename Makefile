bench:
	@go test -bench . > bench.txt
	@cat bench.txt

bench-full:
	@go test -run=. -bench=. -benchtime=5s -count 5 -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out . | tee bench.txt
