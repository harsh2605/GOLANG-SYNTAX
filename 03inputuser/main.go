package main

import (
	"bufio"
	"fmt"
	"os"
)
//comma ok(input,_) or comma err(input , err) syntax in go language
func main() {
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Assign a rating for my pizza")
	input ,_:=reader.ReadString('\n')
	fmt.Println("The current rating of your pizza is",input)
	fmt.Printf("Type of this rating is %T",input)
}