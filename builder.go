package main

import (
	"strings"
)

func StringsBuilder(hello []string) {
	var b strings.Builder

	for _, str := range hello {
		b.WriteString(str)
	}
	// 	fmt.Println(b.String())
}
