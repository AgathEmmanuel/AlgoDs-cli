/*

https://leetcode.com/problems/find-numbers-with-even-number-of-digits/

Results for the solution using normal iteration to find not of digits

- Runtime: 5ms
- Memory: 3.2MB


*/

package main

import (
	"fmt"
	"math"
)

func findNumbers(nums []int) int {
	var count int=0;
	for _,num:=range nums {
		if evenOrOdd(num) {
			count++
		}
	}
	return count
}

func noOfDigits(inputNumber int) int {
	if inputNumber<0 {
		inputNumber=inputNumber*-1
	}
	if inputNumber==0 {
		return 1
	}
	var count int=0;
	for inputNumber>0 {
		count++
		inputNumber=inputNumber/10
	}
	return count
}

func noOfDigitsWithLog(inputNumber int) int {
	if inputNumber<0 {
		inputNumber=inputNumber*-1
	}
	if inputNumber==0 {
		return 1
	}
	var count int=int(math.Log10(float64(inputNumber))+1);
	return count
}
func evenOrOdd(inputNumber int) bool {
	var digitsNum int=noOfDigitsWithLog(inputNumber);
	if digitsNum%2==0 {
		return true
	}
	return false
}


func main(){
	nums:=[]int{1213,1,11,22,333,912}
	fmt.Println(findNumbers(nums))
	fmt.Println(noOfDigits(11122))
	fmt.Println(noOfDigits(-11122))
	fmt.Println(noOfDigits(0))
}
