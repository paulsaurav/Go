package main

import "fmt"

const LoginToken string = "askdnkandabsdbnan"

func main() {
	var username string = "Saurav"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.4646101616016
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var aVariable int
	fmt.Println(aVariable)
	fmt.Printf("Variable is of type: %T \n", aVariable)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//explicit type
	var website = "example.com"
	fmt.Println(website)

	//no var type
	noOfUser := 500000
	fmt.Println(noOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
