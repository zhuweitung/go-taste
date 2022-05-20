// @author zhuweitung 2022/5/20
package main

import (
	"fmt"
	"os"
)

func main() {
	// 从标准输入读取数据
	//scanf()
	//scan()
	//scanln()

	// 从文件读取数据
	//fscanf()

	// 从字符串读取数据
	sscan()
}

// 输入值之间用空格，只能输入一行数据
func scanf() {
	fmt.Println("scanf...")
	var num1 int
	var num2 int
	fmt.Scanf("%d%d", &num1, &num2)
	fmt.Println(num1, num2)
}

// 输入值之间用空格或者换行，不用指定格式化字符串
func scan() {
	fmt.Println("scan...")
	var num1 int
	var num2 int
	fmt.Scan(&num1, &num2)
	fmt.Println(num1, num2)

	var num3 float32
	var num4 float32
	fmt.Scan(&num3, &num4)
	fmt.Println(num3, num4)
}

// 输入值之间用空格，不用指定格式化字符串, 并且只能输入一行数据
func scanln() {
	fmt.Println("scanln...")
	var num1 int
	var num2 int
	fmt.Scanln(&num1, &num2)
	fmt.Println(num1, num2)
}

// 从文件读取数据
func fscanf() {
	var num1 int
	var num2 int
	fmt.Fscanf(os.Stdin, "%d%d", &num1, &num2)
	fmt.Println(num1, num2)
}

// 从字符串读取数据
func sscan() {
	str := "lnj 33"
	var name string
	var age int
	fmt.Sscan(str, &name, &age)
	fmt.Println("name =", name, "age =", age)
}
