package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the current rating of your pizza")
	input ,_:=reader.ReadString('\n')
	fmt.Println("The current rating of your pizza is",input)
	//Suppose you want to add 1 to the rating which is coming but it is of type string and if we are giving input as suppose 56 it is coming as 56\n so with the conversion of string to number you have to remove the \n also

	numrating,err:=strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err!=nil{
		fmt.Println(err)
	}else
	{
		fmt.Println("The current rating of pizza is changed to ",numrating+1)
	}

}