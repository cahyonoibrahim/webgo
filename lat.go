package main

import (
	"fmt"
	"os"
)

func main() {
	j, _ := os.Getwd()
	fmt.Println(j)
}
