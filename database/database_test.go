package database

import "testing"

func BenchmarkGetFoo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Get("foo")
	}
}
