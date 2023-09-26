package main

import "testing"

func BenchmarkStringBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringsBuilder(names)
	}
}
