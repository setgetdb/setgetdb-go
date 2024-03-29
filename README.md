<img src="./logo.png" width="400" />

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
BenchmarkSetDifferentValues-8   	   10000	    100430 ns/op
BenchmarkGetSameValues-8        	100000000	        19.5 ns/op
BenchmarkSetDifferentKeys-8     	   10000	    136741 ns/op
PASS

Process finished with exit code 0
```

## Next steps
- [X] Set operation
- [X] Get operation
- [X] Delete operation
- [X] Server
- [X] Cache layer
- [X] Thread-safe filemanager layer
- [X] Thread-safe cachemanager layer

Gopher image by [Renee French][rf], licensed under [Creative Commons 3.0 Attributions license][cc3-by].

[rf]: https://reneefrench.blogspot.com/
[cc3-by]: https://creativecommons.org/licenses/by/3.0/
