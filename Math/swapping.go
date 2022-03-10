package main

import (
	"fmt"
)

func swap(arr []int,firstIndex int,secondIndex int) []int {
	arr[firstIndex],arr[secondIndex]=arr[secondIndex],arr[firstIndex]
	return arr
}
func swapt(arr []int,firstIndex int,secondIndex int) []int {
	tmp:=arr[firstIndex]
	arr[firstIndex]=arr[secondIndex]
	arr[secondIndex]=tmp
	return arr
}
func swapPtr(arr []int,firstIndex *int,secondIndex *int) []int {
	//fmt.Println(firstIndex)
	//fmt.Println(secondIndex)
	arr[*firstIndex],arr[*secondIndex]=arr[*secondIndex],arr[*firstIndex]
	return arr
}
func main() {
	ex1:=[]int{0,1,2,3,4,5}
	fmt.Println(ex1)
	fmt.Println((&ex1[3]))
	fmt.Println(*(&ex1[3]))
	swap(ex1,3,1)
	fmt.Println(ex1)
	swapPtr(ex1,&ex1[4],&ex1[0])
	fmt.Println(ex1)
	fmt.Println()
	swapt(ex1,4,2)
	fmt.Println(ex1)
}