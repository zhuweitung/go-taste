// @author zhuweitung 2022/5/20
package main

import "fmt"

func main() {
	fmt.Println("calculate result is", calculate(1, 2))
}

// @desc 计算
// @param
// @return
// @author zhuweitung
// @date 2022/5/20
func calculate(numberA int, numberB int) int {
	return numberA + numberB
}
