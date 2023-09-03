package main

import "fmt"

func main() {
	fmt.Println("Here is your description for using copy of array")
	arr :=[]int{45,56,67,78,77,66,55,44,67,78,77,66,55,44,33,22,11,12,23}
	// slicearr :=arr[:len(arr)-10]
	new_slice :=make([]int,len(arr))
	arr=append(arr,45,56,67,78)
	copy(new_slice,arr)
	fmt.Println(new_slice)
	fmt.Println(arr)
}