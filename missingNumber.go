/*

https://leetcode.com/problems/missing-number/

array nums containing n distinct numbers in the range [0, n]

method 1: do a cyclic sort and then iterate to find the missing value
runtime: 16ms
memory: 6.3MB

method 2: exected sum of n numbers= (n*(n+1))/2
		  so missing number= expected sum - actual sum of array
runtime: 12ms
memory: 6.3MB

*/
package main

import (
	"fmt"
)

func missingNumberMethod1(nums []int) int {
	swap:=func(arr []int,i1 int,i2 int) []int {
		arr[i1],arr[i2]=arr[i2],arr[i1]
		return arr
	}
	for i:=0;i<len(nums);{
		correctIndex:=nums[i]
		if (nums[i]<len(nums) && nums[i]!=nums[correctIndex]) {
			swap(nums,i,correctIndex)
		} else {
			i++
		}
	}
	// to find the missing number or the missing index
	for i:=0;i<len(nums);i++ {
		if nums[i]!=i {
			return i
		}
	}
	// if n is not present in index
	return len(nums)
}
func missingNumberMethod2(nums []int) int {
    sum:=func(arr []int) int {
        s:=(len(arr)*(len(arr)+1)/2)
        return s
    }
    suma:=func(arr []int) int {
        sa:=0
        for i:=0;i<len(arr);i++ {
            sa+=arr[i]
        }
        return sa
    }
    mn:=sum(nums)-suma(nums)
    return mn
}


func main() {
	ex1:=[]int{3,0,1}
	ex2:=[]int{0,1}
	ex3:=[]int{9,6,4,2,3,5,7,0,1}
	ex4:=[]int{0}
	fmt.Println(missingNumberMethod1(ex1))
	fmt.Println(missingNumberMethod1(ex2))
	fmt.Println(missingNumberMethod1(ex3))
	fmt.Println(missingNumberMethod1(ex4))


	fmt.Println()
	fmt.Println(missingNumberMethod1(ex1))
	fmt.Println(missingNumberMethod1(ex2))
	fmt.Println(missingNumberMethod1(ex3))
	fmt.Println(missingNumberMethod1(ex4))
}