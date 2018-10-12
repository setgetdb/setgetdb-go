# SetGetDB
A persistent key-value db for educational purposes Only.

## How to use
```go
package main

import (
	"fmt"
	"github.com/setget/setget/src"
)

const TEST = "TEST"
const HELLO = "HELLO"
const WORLD = "WORLD"

func main()  {
    db := setget.NewDatabase(TEST)
    db.Set(HELLO, WORLD)
    _, value := db.Get(HELLO)
    fmt.Println(value)
    db.Delete(HELLO)
    db.Close()
}
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
pkg: github.com/setget/setget/tests
BenchmarkSetDifferentValues-8   	   20000	     80098 ns/op
BenchmarkSetDifferentKeys-8     	   20000	     97435 ns/op
PASS

Process finished with exit code 0
```

## Next steps
- [X] Set operation
- [X] Get operation
- [X] Delete operation
- [ ] Thread-safe
- [ ] Server
