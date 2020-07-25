// +build ignore

package main

import (
	"fmt"
	"github.com/ikaz7/squash"
)

func main() {
	s := []byte{'a', ' ', ' ', 'd', 'e', 'f', 'f', ' ', ' ', ' ', 'f'}
	fmt.Printf("%s\n", s)
	s = squash.Space(s)
	fmt.Printf("%s\n", s)
}
