package main

import (
	"fmt"
	"os"
)

type CustomReader struct{}

func (CustomReader) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Error: missing the file name")
		os.Exit(1)
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	// v1 - Logs the file data
	// io.Copy(CustomReader{}, file)

	// v2 - Logs the file data
	// fmt.Println()
	// data, err := io.ReadAll(file)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(data))

	// v3 - Logs the file data
	fmt.Println()
	bs := make([]byte, 32*1024)
	totalBytes, err := file.Read(bs)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Total bytes:", totalBytes)
	fmt.Println(string(bs))
}
