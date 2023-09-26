package main

import "testing"

func BenchmarkStringBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringBuffer(names)
	}
}
