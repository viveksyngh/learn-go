package main

import "os"

func main() {
	panic("Got a problem")
	_, err := os.Create("/tmp/testfile")

	if err != nil {
		panic(err)
	}
}
