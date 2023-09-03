package main

import "fmt"

func main() {
	var arr [5]string
	arr[0] = "harsh"
	arr[4] = "aniket"
	fmt.Println(arr)

	var second_arr = [5]string{}
	//second_arr[0]="harsh"
	fmt.Println(len(second_arr))
}