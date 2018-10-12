# SetGetDB

<img src="./logo.png" width="600" />

A persistent key-value db for educational purposes only.

## How to build
```sh
$ make build
```

## How to run
```sh
$ docker run -d -p 10101:10101 setgetdb/setgetdb
```

## How to use

### Set a value
```sh
$ curl -X POST --data '{ "key": "hello", "value": "world" }' http://localhost:10101/set
```

### Retrieve a value
```sh
$ curl -X POST --data '{ "key": "hello" }' http://localhost:10101/get
```

### Delete a value
```sh
$ curl -X POST --data '{ "key": "hello" }' http://localhost:10101/delete
```

## Benchmark

### Hardware
- **Processor**: Intel(R) Core(TM) i7-4850HQ CPU @ 2.30GHz
- **RAM**: 16GB
- **Disk**: SSD 500GB

### Results
```
goos: darwin
goarch: amd64
pkg: github.com/setgetdb/setgetdb/tests
BenchmarkSetDifferentValues-8   	   20000	     79983 ns/op
BenchmarkGetSameValues-8        	20000000	        67.5 ns/op
BenchmarkSetDifferentKeys-8     	   20000	    144688 ns/op
PASS

Process finished with exit code 0
```

## Next steps
- [X] Set operation
- [X] Get operation
- [X] Delete operation
- [X] Server
- [X] Cache layer
- [ ] Thread-safe
- [ ] Cache LRU strategy

