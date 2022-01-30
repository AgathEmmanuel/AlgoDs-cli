/*


Find a targetElement in an array


Array should be a sorted array
If sorted in increasing order [ 1,3,9...,10]

steps:

1. Find the middleElement ((startOfArray + endOfArray) / 2)
2. if targetElement > middleElement: 
	Search in right part of array
	startOfArray = middleElement+1
   else: 
   	Search in left part of array
	endOfArray = middleElement-1
3. if middleElement == targetElement: 
	answer is middleElement
4. if startOfArray > endOfArray: 
	element not exist


Complexity anlaysis:

Best case:  
	 when middleElement == targetElement in the first iteration
	 input size of array do not matter
	 O(1)
     time|
	 |_____
	 |_______ 
	     size of input 

Worst case:
	 when middleElement == targetElement in the last iteration
	 at level 1   : input size of array is N
	 at level 2   : N becomes N/2
	 at level 3   : N becomes N/4
	 at level 4   : N becomes N/8
	 at level (k) : N becomes N/2^k
	 		1 = N/2^k
			N = 2^k
			log(N) = log(2^k)
			log(N) = k*log(2)
			k = log(N) / log(2)
			  = log2(N)     [ the base constants dont matter]
 	So the total no of levels or iterations in worst case is log2(N)

	ex: N = 1000000
	    k = log2(N) = 20
	    Binary search : no of comparisons = 20 
	    Linear search : no of comparisons = 1000000



If sorted in decreasing order [ 10,9,3...,1]

if targetElement > middleElement:
	search left
	endOfArray = targetElement-1
else:
	search right
	startOfArray = targetElement+1



Order Agnostic Binary Search:

if startOfArray > endOfArray : decreasing order sorted
else : increasing order sorted


*/




package main

import (
	"fmt"
)

func orderAgnosticBinarySearch(inputArray []int,targetElement int) int {
	startOfArray:=0
	endOfArray:=len(inputArray)-1
	// find array is ascending or descending order sorted
	var isAscending bool=inputArray[startOfArray] < inputArray[endOfArray]
	/*
	if inputArray[startOfArray] < inputArray[endOfArray] {
		isAscending=true
	} else {
		isAscending=false
	}
	*/
	for startOfArray < endOfArray {

		fmt.Println("iteration")
		middleIndex := startOfArray + ((endOfArray - startOfArray) / 2)
		if inputArray[middleIndex]==targetElement {
			return middleIndex
		}
		if isAscending {
			if targetElement < inputArray[middleIndex] {
				endOfArray = middleIndex-1
			} else {
				startOfArray = middleIndex+1
			}
		} else {
			if targetElement < inputArray[middleIndex] {
				startOfArray = middleIndex+1
			} else {
				endOfArray = middleIndex-1
			}
		}
	}
	return -1
}



func binarySearch(inputArray []int,targetElement int) int {
	// return -1 if targetElement not exist
	startOfArray:=0
	endOfArray:=len(inputArray)-1
	// find the middleElement
	// middleElement = ((startOfArray + endOfArray) / 2)
	// there can be cases when (startOfArray + endOfArray) exceeds the range of integer available 
	// so we can use this technique
	// middleIndex = startOfArray + ((endOfArray - startOfArray) / 2)
	// m = ((s+e)/2)
	//   = (s/2) + (e/2)
	//   = (s/2) + (e/2) + ( (s/2) - (s/2) )
	//   = s + (e/2) - (s/2) 
	//   = s + (e - s)/2 

	for startOfArray < endOfArray {

		fmt.Println("iteration")
		middleIndex := startOfArray + ((endOfArray - startOfArray) / 2)
		if targetElement < inputArray[middleIndex] {
			endOfArray = middleIndex-1
		} else if targetElement > inputArray[middleIndex] {
			startOfArray = middleIndex+1
		} else {
			return middleIndex
		}
	}
	return -1 
}


func main() {
	ex1:=[]int{-17,-11,-9,2,21,22,30,35,40,41}
	ex2:=[]int{17,11,9,2,1,-22,-30,-35,-40,-41}
	fmt.Printf("%T\n",ex1)
	fmt.Println(ex1)
	var targetElement1 int=-9
	var targetElement2 int=-35
	//var targetIndex int=binarySearch(ex1,targetElement)
	var targetIndex1 int=orderAgnosticBinarySearch(ex1,targetElement1)
	var targetIndex2 int=orderAgnosticBinarySearch(ex2,targetElement2)
	fmt.Println("targetElement",targetElement1)
	fmt.Println(targetIndex1)
	fmt.Println("targetElement",targetElement2)
	fmt.Println(ex2)
	fmt.Println(targetIndex2)
}

