package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating you want to give: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating!!!, ", input)

}
