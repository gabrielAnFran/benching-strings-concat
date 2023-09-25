package main

import (
	"bytes"
	"fmt"
)

func StringBuffer(hello []string) {
	var b bytes.Buffer


	for _, str:= range hello {
		b.WriteString(str)
	}

	fmt.Println(b.String())
}
