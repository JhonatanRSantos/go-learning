package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type CustomWriter struct{}

func (CustomWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("https://www.google.com.br/")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// manual way
	// 1 MB of size
	// bs := make([]byte, 1000000)

	// totalBytes, err := resp.Body.Read(bs)
	// defer resp.Body.Close()

	// if err != nil {
	// 	fmt.Println("Error:", err.Error())
	// }
	// fmt.Println("Total of Bytes", totalBytes)
	// fmt.Println(string(bs))

	// Prints all data into console
	// io.Copy(os.Stdout, resp.Body)

	// customWriter := CustomWriter{}

	// io.Copy(customWriter, resp.Body)

	data, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(data))
}
