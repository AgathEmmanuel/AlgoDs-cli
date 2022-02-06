
/*


*/

package main

import (
	"fmt"
	"math"
)


func nthFibFormula (n float64) float64 {
	nthFibNum:=math.Pow(((1+math.Sqrt(5))/2),n)/math.Sqrt(5)
	return nthFibNum
}
func main() {
	fmt.Println("Fibonacci number using golden ration")
	var i float64=0
	for i=0;i<10;i++ {
		//fmt.Println(nthFibFormula(i))
		fmt.Println(math.Pow(i,2))
	}
}
