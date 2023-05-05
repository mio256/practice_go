package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	s := time.Now()

	var builder strings.Builder

	builder.Grow(10000)

	builder.WriteByte('a')

	for i := 0; i < 10000; i++ {
		builder.WriteByte('b')
	}

	fmt.Printf("%s\n", builder.String())
	fmt.Printf("process time: %s\n", time.Since(s))
}
