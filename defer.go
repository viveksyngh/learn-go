package main

import "fmt"
import "os"

func main() {
	file := createFile("/tmp/defer.txt")
	defer closeFile(file)
	writeFile(file)
}

func createFile(p string) *os.File {
	fmt.Println("Creating file")
	file, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return file
}

func writeFile(file *os.File) {
	fmt.Println("Writing to file")
	fmt.Fprintln(file, "Data")
}

func closeFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}
