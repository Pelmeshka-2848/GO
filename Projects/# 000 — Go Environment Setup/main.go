package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go environment is ready")
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
}
