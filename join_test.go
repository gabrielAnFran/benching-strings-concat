package main

import "testing"

func BenchmarkJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Join(names)
	}
}
