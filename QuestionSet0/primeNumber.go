/*

Only make checks for numbers less than the square root of the number iself  
then the complexity will be Sqrt(N)


*/

package main

import (
	"fmt"
)


func isPrime (n int) bool {
	if n <= 1 {
		return false
	}
	// we are checking till square root only since after that its 
	// just getting repeated only the arrangement will be changing  
	if b:=2;b*b<=n {
		if n%b==0 {
			return false
		}
		b+=1
	}
	return true
	// time complexity: Sqrt(n)
}


func arrayFilter(n int, primes [51]bool) {
	// https://www.geeksforgeeks.org/sieve-of-eratosthenes/
	// sieve-of-eratosthenes method
	// 
	// eliminating elements from the array having numbers till n, based on wheather
	// they are multiples of the previous elements
	// while iterating throgh the array all the multiples of the corresponding primes till
	// square root of n will be removed hence resulting in an array with only prime numbers
	// till n
	fmt.Println(primes)
	fmt.Printf("%T \n",primes)
	for i:=2;i*i<=n;i++ {
		if primes[i]==false {
			for j:=i*2;j<=n;j+=i {
				primes[j]=true
			}
		}
	}
	for i:=2;i<=n;i++ {
		if primes[i]==false {
			fmt.Println(i," ")
		}
	}
	// Time Complexity: 
	// no of numbers till n that can be devided by 2 = n/2
	// no of numbers till n that can be devided by 5 = n/5
	// n/2 + n/5 + n/7 + ..... + n/p
	// the value you can calculate using harmonic progression
	// n*( 1/2 + 1/5 + 1/7 + .... + 1/p )
	// using Taylor series  
	// n*( log(log(n)) )



}
func main() {
	fmt.Println("Prime number ")
	n:=20
	for i:=1;i<=n;i++ {
		fmt.Println(i," is ",isPrime(i))
	}
	m:=50
	var primes [50+1]bool
	arrayFilter(m,primes)
}
