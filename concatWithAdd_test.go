package main

import "testing"

func BenchmarkConcatWithAdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ConcatWithAdd(names)
	}
}
