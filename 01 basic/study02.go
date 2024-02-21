package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unsafe"
)

func main() {
	var i8 int8
	var i16 int16

	fmt.Print(i8, unsafe.Sizeof(i8), math.MaxInt8, math.MinInt8, "\n")
	fmt.Print(i16, unsafe.Sizeof(i16), math.MaxInt16, math.MinInt16, "\n")

	//字符串拼接
	name := "tom"
	native := "合肥"
	age := 21

	msg1 := name + native
	msg2 := fmt.Sprint(name, age)
	msg3 := strings.Join([]string{name, native}, "")

	var buffer bytes.Buffer
	buffer.WriteString(name)
	buffer.WriteString(strconv.Itoa(age)) //强转
	buffer.WriteString(native)

	fmt.Print(msg1, "\n")
	fmt.Print(msg2, "\n")
	fmt.Print(msg3, "\n")
	fmt.Println(buffer.String())

	fmt.Println(len(msg3))
}
