package main

import (
	"fmt"
	"os"
)


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing pparameter, provide filename!")
		return
	}

	data, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}

	fmt.Println("FIle content is")
	fmt.Println(string(data))

}