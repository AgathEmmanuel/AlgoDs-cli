/*
Cycle Sort

in-place algorithm
a comparison sort
minimizes no of memory writes

--- logic
if values position is not correct
	written(swapped) one time to its correct position
else
    no writting(swaping) needed


*/
package main

import (
	"fmt"
)

func swap(arr []int,firstIndex int,secondIndex int) []int {
	arr[firstIndex],arr[secondIndex]=arr[secondIndex],arr[firstIndex]
	return arr
}

func cyclicSort(inputArray []int) {
	for i:=0;i<len(inputArray);{
		correctIndex:=inputArray[i]-1
		if inputArray[correctIndex]!=inputArray[i] {
			swap(inputArray,i,correctIndex)
		} else {
			i++;
		}
	}
}
func main() {
	ex1:=[]int{3,1,4,6,5,2}
	fmt.Println(ex1)
	cyclicSort(ex1)
	fmt.Println(ex1)
}