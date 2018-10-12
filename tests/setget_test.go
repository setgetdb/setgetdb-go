package tests

import (
	. "github.com/setgetdb/setgetdb/setgetdb"
	"testing"
)

const PREFIXPATH = "./"
const TEST = "TEST"
const HELLO = "HELLO"
const WORLD = "WORLD"

func BenchmarkSetDifferentValues(b *testing.B) {
	db := NewDatabase(PREFIXPATH, TEST)
	for n := 0; n < b.N; n++ {
		db.Set(HELLO, WORLD + string(n))
	}
}

func BenchmarkGetSameValues(b *testing.B) {
	db := NewDatabase(PREFIXPATH, TEST)
	db.Set(HELLO, WORLD)
	for n := 0; n < b.N; n++ {
		db.Get(HELLO)
	}
}

func BenchmarkSetDifferentKeys(b *testing.B) {
	db := NewDatabase(PREFIXPATH, TEST)
	for n := 0; n < b.N; n++ {
		db.Set(HELLO + string(n), WORLD)
	}
}