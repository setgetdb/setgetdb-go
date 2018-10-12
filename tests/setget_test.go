package tests

import (
	. "github.com/setgetdb/setgetdb/setgetdb"
	"testing"
)

const TEST = "TEST"
const HELLO = "HELLO"
const WORLD = "WORLD"

func BenchmarkSetDifferentValues(b *testing.B) {
	db := NewDatabase(TEST)
	for n := 0; n < b.N; n++ {
		db.Set(HELLO, WORLD+string(n))
	}
}

func BenchmarkSetDifferentKeys(b *testing.B) {
	db := NewDatabase(TEST)
	for n := 0; n < b.N; n++ {
		db.Set(HELLO+string(n), WORLD)
	}
}