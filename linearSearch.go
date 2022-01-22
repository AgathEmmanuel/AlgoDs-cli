/*


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

func main() {
	ex1:=[]int{9,3,8,1,2,3,4,0,-2,22}
//	ex1:=[]int{}
	fmt.Println(ex1)
//	var element int=999
	var element int=-2
	var output int=linearSearch(ex1,element);
	fmt.Println(output)
}
