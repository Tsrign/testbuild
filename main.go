package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("hello")
	fmt.Println(getCurrentFilePath())
	fmt.Println("success")
}

func getCurrentFilePath() string {
	_, filePath, _, _ := runtime.Caller(1)
	return filePath
}
