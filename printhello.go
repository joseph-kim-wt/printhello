package main

import (
	"fmt"

	"github.com/joseph-kim-wt/sayhello"
)

func main() {
	message := sayhello.Sayhello("joseph")
	fmt.Println(message)
}
