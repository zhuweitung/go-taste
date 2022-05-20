// @author zhuweitung 2022/5/20
package main

import (
	"fmt"
)

func main() {
	// 0110
	numberA := 6
	// 1011
	numberB := 11
	// &^为逻辑清零运算符, B对应位是1,A对应位清零,B对应位是0, A对应位保留原样
	result := numberA &^ numberB
	// 0100
	fmt.Println(result)
}
