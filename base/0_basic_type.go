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

}
