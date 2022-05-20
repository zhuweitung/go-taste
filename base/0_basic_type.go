// @author zhuweitung 2022/5/20
package main

import (
	"fmt"
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
}

// 了解各种类型的默认值
func defaultValue() {
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
