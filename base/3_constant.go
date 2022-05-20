// @author zhuweitung 2022/5/20
package main

import "fmt"

const PI float64 = 3.1415926

//const PI = 3.1415926 // 也可

// 多重赋值方式
const MsgTitle, MsgTemplate = "消息提醒", "%s\n  你好，%s"

// 常量组方式
const (
	ApiUrlGet    = "127.0.0.1:8080/proc/get/1" // 常量未被使用，编译不会报错
	ApiUrlDelete = "127.0.0.1:8080/proc/delete/1"
)

// 在常量组中, 如果上一行常量有初始值,但是下一行没有初始值, 那么下一行的值就是上一行的值
const (
	MsgValue1 = "111"
	MsgValue2 // "111"
)

// 枚举常量
// 利用iota标识符标识符实现从0开始递增的枚举
const (
	male   = iota
	female = iota
)

// 在同一个常量组中,只要上一行出现了iota,那么后续行就会自动递增
//const (
//	male = iota // 这里出现了iota
//	female // 这里会自动递增
//)

func main() {
	//PI = 1 // 报错，常量不可再变更
	fmt.Println("PI =", PI)

	fmt.Println(fmt.Sprintf(MsgTemplate, MsgTitle, "go"))

	fmt.Println(MsgValue1, MsgValue2)

	fmt.Println(male, female)
}
