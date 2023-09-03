package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = []string{"harsh", "yash", "raj"}
	arr = append(arr, "samal", "jijnansu")
	arr=append(arr[1:1])//Here the last range is not inclusive
	fmt.Println(arr)

	arr2:=make([]int,4)
	arr2[0]=2
	arr2[1]=4
	arr2=append(arr2,45)
	// arr2=append(arr2,12,23,34,45,56,67,78)
	fmt.Println(len(arr2))
	fmt.Println(cap(arr2))
	sort.Ints(arr2)
	fmt.Println(arr2)

}