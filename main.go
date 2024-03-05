package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	TEST_DIRECTORY = "testing_files"
)

func main() {
	fmt.Printf("Hello,world!\n")
	file, err := os.Open(TEST_DIRECTORY + "/" + "one.json")
	if err != nil {
		fmt.Println("err : ", err)
		os.Exit(1)
	}

	var emptyInterface interface{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&emptyInterface)

	if err != nil {
		fmt.Println("err : ", err)
		os.Exit(1)
	}
	fmt.Println("Interface -> ", emptyInterface)
}
