package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x := 5
	bytes, err := json.Marshal(x)
	if err != nil {
		fmt.Println("Can't serislize", x)
	}

	fmt.Printf("%v => %v, '%v'\n", x, bytes, string(bytes))

}
