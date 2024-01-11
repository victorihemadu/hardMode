package main

import (
	"fmt"
	"io"
	"os"
)


func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Missing pparameter, provide filename!")
	// 	return
	// }

	// data, err := os.ReadFile(os.Args[1])

	// if err != nil {
	// 	fmt.Println("Can't read file:", os.Args[1])
	// 	panic(err)
	// }

	// fmt.Println("FIle content is")
	// fmt.Println(string(data))

	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)

}