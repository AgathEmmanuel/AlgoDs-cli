/*



*/



package main

import (
	"fmt"
)


func binarySearchSqrt(n int, precision int)float64 {
	start:=0
	end:=n
	root := 0.0

	for start=0;start<=end; {
		middle:= start+(end-start)/2
		if middle*middle == n {
			return float64(middle)
		}
		if middle*middle > n {
			end = middle-1
		}else {
			start = middle+1
		}
	}
	increment:=0.1
	for i:=0;i<=precision;i++ {
		for root*root <=float64(n) {
			root += increment
		}
		root-=increment
		increment/=10
	}
	return root
}





func main() {
	fmt.Println("finding sqrt of a number")
	n:=50
	precision:=5
	fmt.Println("The sqrt is ",binarySearchSqrt(n,precision))
}

