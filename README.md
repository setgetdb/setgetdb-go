# SetGetDB
A persistent key-value db for educational purposes Only.

## How to use
```go
const TEST = "TEST"
const HELLO = "HELLO"
const WORLD = "WORLD"

db := setget.NewDatabase(TEST)
db.SetByKey(HELLO, WORLD)
value := db.GetValueByKey(HELLO)
fmt.Println(value)
db.DeleteByKey(HELLO)
```

## Next steps
- [X] Set operation
- [X] Get operation
- [X] Delete operation
- [] Thread-safe
- [] Server
