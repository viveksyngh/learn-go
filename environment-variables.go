package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR", os.Getenv("BAR"))

	fmt.Println()
	for _, env := range os.Environ() {
		pair := strings.Split(env, "=")
		fmt.Println(pair[0], ":", pair[1])
	}
}