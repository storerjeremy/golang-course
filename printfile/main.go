package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Something bad happened")
	}
	b, _ := ioutil.ReadAll(file)
	fmt.Println(string(b))
}
