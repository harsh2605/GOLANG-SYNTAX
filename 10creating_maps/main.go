package main

import "fmt"

func main() {
	fmt.Println("Using map in this file in a fun way")
	//Syntax is similar to that of using make function for making slices
	arr := make(map[string]string)//The string inside the square bracket is called key and the outside one is called value
	arr["harsh"]="loves to learn GOLANG"
	arr["aniket"]="loves to learn GOLANG"
	arr["subham"]="loves to learn GOLANG"
	 fmt.Println(arr)

	// //For deleting any key value from the map
	// delete(arr,"aniket")
	// fmt.Println(arr["harsh"])

	//Iterating over the maps using for each loop
	for key,value := range arr{
		fmt.Printf("For the key the value is %v and for the value the value is %v\n",key,value)
	}
	//If you dont want to use any of the two values key or values use COMMAOK syntax
	for _,value := range arr{
		fmt.Printf("For the value the value is %v\n ",value)
		if(8>9){
			fmt.Println("yes");
		}else
		{
			fmt.Println("No")
		}
	}
}