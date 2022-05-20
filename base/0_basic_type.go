// @author zhuweitung 2022/5/20
package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var intNum int = 1
	fmt.Println("int size is", unsafe.Sizeof(intNum))

	var int8Num int8 = 2
	fmt.Println("int8 size is", unsafe.Sizeof(int8Num))

	var int16Num int16 = 3
	fmt.Println("int16 size is", unsafe.Sizeof(int16Num))

	var int32Num int32 = 4
	fmt.Println("int32 size is", unsafe.Sizeof(int32Num))

	var int64Num int64 = 5
	fmt.Println("int64 size is", unsafe.Sizeof(int64Num))

	var uintptrNum uintptr = 5
	fmt.Println("uintptr size is", unsafe.Sizeof(uintptrNum))

	var uint8Num uint8 = 5
	fmt.Println("uint8 size is", unsafe.Sizeof(uint8Num))

	var uint16Num uint16 = 5
	fmt.Println("uint16 size is", unsafe.Sizeof(uint16Num))

	var uint32Num uint32 = 5
	fmt.Println("uint32 size is", unsafe.Sizeof(uint32Num))

	var uint64Num uint64 = 5
	fmt.Println("uint64 size is", unsafe.Sizeof(uint64Num))

	var byteChar byte = 'a'
	fmt.Println("byte size is", unsafe.Sizeof(byteChar))

	var runeNum rune = 1
	fmt.Println("rune size is", unsafe.Sizeof(runeNum))

	var float32Num float32 = 0.1
	fmt.Println("float32 size is", unsafe.Sizeof(float32Num))

	var float64Num float64 = 0.1
	fmt.Println("float64 size is", unsafe.Sizeof(float64Num))

	// 多变量连续定义
	numberA, numberB := 1, 2
	fmt.Println(numberA, numberB)

	// 变量组
	var (
		numberC, numberD = 3, 4
		numberE, numberF = 9.99, 100
	)
	fmt.Println(numberC, numberD, numberE, numberF)

	defaultValue()
	convert()
	convertToString()
	stringToX()
}

// 了解各种类型的默认值
func defaultValue() {
	fmt.Println("defaultValue...")
	var intV int                 // 整型变量
	var floatV float32           // 实型变量
	var boolV bool               // 布尔型变量
	var stringV string           // 字符串变量
	var pointerV *int            // 指针变量
	var funcV func(int, int) int // function变量
	var interfaceV interface{}   // 接口变量
	var sliceV []int             // 切片变量
	var channelV chan int        // channel变量
	var mapV map[string]string   // map变量
	var errorV error             // error变量

	fmt.Println("int = ", intV)                  // 0
	fmt.Println("float = ", floatV)              // 0
	fmt.Println("bool = ", boolV)                // false
	fmt.Println("string = ", stringV)            // ""
	fmt.Println("pointer = ", pointerV)          // nil
	fmt.Println("func = ", funcV)                // nil
	fmt.Println("interface = ", interfaceV)      // nil
	fmt.Println("slice = ", sliceV)              // []
	fmt.Println("slice = nil is", sliceV == nil) // true
	fmt.Println("channel = ", channelV)          // nil
	fmt.Println("map = ", mapV)                  // map[]
	fmt.Println("map = nil is", mapV == nil)     // true
	fmt.Println("error = ", errorV)              // nil

	var arraryV [3]int // 数组变量
	type Person struct {
		name string
		age  int
	}
	var structV Person                // 结构体变量
	fmt.Println("arrary = ", arraryV) // [0, 0, 0]
	fmt.Println("struct = ", structV) // {"" 0}
}

// 数据类型转换
func convert() {
	fmt.Println("convert...")
	var num0 int = 10
	var num1 int8 = 20
	var num2 int16
	//num2 = num0 // 编译报错, 不同长度的int之间也需要显示转换
	num2 = int16(num0)
	fmt.Println(num0, num1, num2)

	var num3 float32 = 3.14
	var num4 float64
	//num4 = num3 // 编译报错, 不同长度的float之间也需要显示转换
	num4 = float64(num3)
	fmt.Println(num3, num4)

	var num5 byte = 11
	var num6 uint8
	num6 = num5 // 这里不是隐式转换, 不报错的原因是byte的本质就是uint8
	fmt.Println(num5, num6)

	var num7 rune = 11
	var num8 int32
	num8 = num7 // 这里不是隐式转换, 不报错的原因是rune的本质就是int32
	fmt.Println(num7, num8)

}

// 转换为字符串类型
func convertToString() {
	fmt.Println("convertToString...")
	var num1 int32 = 65
	// 可以将整型强制转换, 但是会按照ASCII码表来转换
	var str1 string = string(num1)
	fmt.Println(num1, str1)

	//var num2 float32 = 3.14
	//var str2 string = string(num2) // 编译报错，不能将其它基本类型强制转换为字符串类型
	//fmt.Println(num2, str2)

	//var str3 string = "97"
	//var num3  int = int(str3) // 编译报错，不能强制转换, cannot convert str2 (type string) to type int
	//fmt.Println(num3)

	// 数字类型转字符串
	num3 := 10
	// 第一个参数: 需要被转换的整型,必须是int64类型
	// 第二个参数: 转换为几进制,  必须在2到36之间
	str4 := strconv.FormatInt(int64(num3), 10)
	fmt.Println(num3, str4)

	str5 := strconv.FormatInt(int64(num3), 2)
	fmt.Println(str5)

	var num4 float64 = 3.1234567890123456789
	// 第一个参数: 需要转换的实型, 必须是float64类型
	// 第二个参数: 转换为什么格式,f小数格式, e指数格式
	// 第三个参数: 转换之后保留多少位小数, 传入-1按照指定类型有效位保留
	// 第四个参数: 被转换数据的实际位数,float32就传32, float64就传64
	// 将float64位实型,按照小数格式并保留默认有效位转换为字符串
	str6 := strconv.FormatFloat(num4, 'f', -1, 32)
	fmt.Println(str6) // 3.1234567
	str7 := strconv.FormatFloat(num4, 'f', -1, 64)
	fmt.Println(str7)
	str8 := strconv.FormatFloat(num4, 'f', 2, 64)
	fmt.Println(str8)

	// 布尔类型转字符串
	var num5 bool = true
	str9 := strconv.FormatBool(num5)
	fmt.Println(str9) // true
}

// 字符串类型转换为其他类型
func stringToX() {
	fmt.Println("stringToX...")

	str1 := "127"
	num1, err := strconv.ParseInt(str1, 10, 8)
	fmt.Println(str1, num1, err)

	str2 := "128"
	num2, err := strconv.ParseInt(str2, 10, 8) // int8:-128~127 数据超出范围
	fmt.Println(str2, num2, err)

	str3 := "3.1234567890123456789"
	num3, err := strconv.ParseFloat(str3, 32)
	fmt.Println(str3, num3, err)

	str4 := "false"
	num4, err := strconv.ParseBool(str4)
	fmt.Println(str4, num4, err)

	// 字符串类型转换为数值类型时,如果不能转换除了返回error以外,还会返回对应类型的默认值
	str5 := "abc"
	num5, err := strconv.ParseInt(str5, 10, 32)
	fmt.Println(str5, num5, err) // 0

	// 字符串类型和整型快速转换
	num6 := 123
	str6 := strconv.Itoa(num6)
	fmt.Println(num6, str6)

	str7 := "129"
	num7, err := strconv.Atoi(str7)
	fmt.Println(str7, num7, err)

	// 数据格式化
	num8 := 110
	str8 := fmt.Sprintf("%d", num8)
	fmt.Println(num8, str8)
}
