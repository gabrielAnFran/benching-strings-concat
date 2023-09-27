package main

import "fmt"

func Sprintf(hello []string) {
	var a string

	for _, str := range hello {
		a = fmt.Sprintf(a, str)
	}

}
