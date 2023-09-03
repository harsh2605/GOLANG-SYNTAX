package main

import "fmt"

func main() {
	var arr = []int{34, 45, 56, 67, 78, 89}
	arr = append(arr[:2], arr[3:]...)//With this merhod you can remove any particular index of the arrray of elements present in your current array
	fmt.Println(arr)
}