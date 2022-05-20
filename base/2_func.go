// @author zhuweitung 2022/5/20
package main

import "fmt"

func main() {
	fmt.Println("calculate result is", calculate(1, 2))
	total, result := multipleResult(3, 4)
	fmt.Println("total is", total)
	fmt.Println("result: ", result)
}

// @desc 计算
// @param
// @return
// @author zhuweitung
// @date 2022/5/20
func calculate(numberA int, numberB int) int {
	return numberA + numberB
}

// @desc 多返回值
// @param
// @return
// @author zhuweitung
// @date 2022/5/20
func multipleResult(numberA int, numberB int) (int, string) {
	total := calculate(numberA, numberB)
	return total, fmt.Sprintf("total = %d", total)
}
