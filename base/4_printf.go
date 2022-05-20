// @author zhuweitung 2022/5/20
package main

import "fmt"

func main() {
	// 格式化打印
	num := 16
	// 十进制
	fmt.Printf("%d\n", num)
	// 八进制
	fmt.Printf("%o\n", num)
	// 十六进制
	fmt.Printf("%x\n", num)
	// 二进制
	fmt.Printf("%b\n", num)

	// %T 为数值类型
	num2 := true
	fmt.Printf("the type of num2 is %T\n", num2)
	num3 := 0.111
	fmt.Printf("the type of num3 is %T\n", num3)

	// 任意类型占位符 %v
	fmt.Printf("num=%v, num2=%v, num3=%v\n", num, num2, num3)
}
