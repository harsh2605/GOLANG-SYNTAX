package main

// := this is called walarus operator in go language(only inside the method we are allowed to use this operator)
import "fmt"
func main()  {
	var username string = "harsh"
	fmt.Println(username)
	fmt.Printf("The type of variable username is %T \n",username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("The type of variable username is %T",isLoggedin)

	var website="harsh agawral"
	fmt.Println(website)
	fmt.Printf("The type of variable username is %T",website)

	new_website := 809982
	fmt.Println(new_website)
}