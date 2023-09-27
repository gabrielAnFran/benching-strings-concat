package main

import (
	"testing"
)

func BenchmarkSprintf(b *testing.B) {

	for n := 0; n < b.N; n++ {
		Sprintf(names)
	}

}
