package database

import "testing"

func BenchmarkGetFoo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Get("foo")
	}
}

func BenchmarkCreateFoo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Create([]byte("{ \"foo\": 1234 }"))
	}
}
