package main

import (
	"fmt"
)

func ConcatWithAdd(hello []string) {
	var a string

	for _, str := range hello {
		a = a + str
	}
	fmt.Println(a)
}
