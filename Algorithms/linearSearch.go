/*

max comparisons : N (no of elements)



Time complexity:  O(N)

*/


package main

import (
	"fmt"
)

func linearSearch(inputArray []int,targetElement int) int {
	// if target element not found return -1
	if len(inputArray)==0 {
		return -1;
	}
	// run a for loop to find target element
	for i:=0;i<len(inputArray);i++ {
		if targetElement==inputArray[i] {
			fmt.Printf("The index of element %v is: ",targetElement)
			return i;
		}
	}
	// if target element not found in any of the above cases
	return -1;
}

func linearSearchInRange(inputArray []int,targetElement int, start int, stop int) int {
	// if target element not found return -1
	if len(inputArray)==0 {
		return -1;
	}
	// run a for loop to find target element
	for i:=start;i<stop;i++ {
		if targetElement==inputArray[i] {
			fmt.Printf("The index of element %v is: ",targetElement)
			return i;
		}
	}
	// if target element not found in any of the above cases
	return -1;
}

func smallestElement(inputArray []int) int {
	smallest:=inputArray[0]
	for i:=0;i<len(inputArray);i++ {
		if inputArray[i]<smallest {
			smallest=inputArray[i]
		}
	}
	return smallest;
}

func findIn2DArray(inputArray [][]int,targetElement int) []int {
	for row:=0;row<len(inputArray);row++ {
		for column:=0;column<len(inputArray[row]);column++ {
			if inputArray[row][column]==targetElement {
				return []int{row,column}
			}
		}
	}
	return []int{-1,-1}
}



func linearSearchString(inputString string,targetElement byte) int {
	// if target element not found return -1
	if len(inputString)==0 {
		return -1;
	}
	// run a for loop to find target element
	for i:=0;i<len(inputString);i++ {
		if targetElement==inputString[i] {
			fmt.Printf("The index of element %c is: ",targetElement)
			return i;
		}
	}
	// if target element not found in any of the above cases
	return -1;
}


func main() {
	ex1:=[]int{9,3,8,1,2,3,4,0,-2,22}
//	ex1:=[]int{}
	fmt.Println(ex1)
//	var element int=999
	var element1 int=-2
	var output1 int=linearSearch(ex1,element1);
	fmt.Println(output1)

	ex3:=[]int{9,3,8,1,2,-2,4,0,-2,22,-3,1}
//	ex1:=[]int{}
	fmt.Println(ex3)
//	var element int=999
	var element3 int=-2
	var output3 int=linearSearchInRange(ex3,element3,2,6);
	fmt.Println(output3)

	var output4 int=smallestElement(ex3);
	fmt.Println("The smallest element is:",output4)


	ex5:=[][]int{
		{1,2,9,8,5,4},
		{11,22,1,4,9,91},
		{-1,222,111,0,4},
	}
	var element5 int=-11
	var output5 []int=findIn2DArray(ex5,element5)
	fmt.Println(ex5)
	fmt.Printf("The position of element %d is at ",element5)
	fmt.Println(output5)

	var ex2 string="helllllllabit"
//	ex1:=[]int{}
	fmt.Println(ex2)
//	var element int=999
	var element2 byte='a'
	var output2 int=linearSearchString(ex2,element2);
	fmt.Println(output2)


}
