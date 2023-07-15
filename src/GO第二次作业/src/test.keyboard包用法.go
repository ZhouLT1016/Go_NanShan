package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Println("Press 'Q' to quit.")
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEsc {
			break
		}
		if char == 'q' || char == 'Q' {
			break
		}
		fmt.Printf("You pressed: %q\n", char)
	}
}
