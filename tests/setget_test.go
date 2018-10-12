package tests

import (
	"gitlab.com/setget/setget/src"
	"testing"
)

const TEST = "TEST"
const HELLO = "HELLO"
const WORLD = "WORLD"

func BenchmarkSetDifferentValues(b *testing.B) {
	db := setget.NewDatabase(TEST)
	for n := 0; n < b.N; n++ {
		db.SetByKey(HELLO, WORLD+string(n))
	}
}

func BenchmarkSetDifferentKeys(b *testing.B) {
	db := setget.NewDatabase(TEST)
	for n := 0; n < b.N; n++ {
		db.SetByKey(HELLO+string(n), WORLD)
	}
}