package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	// bool (logic)
	var myBoolean bool = true
	fmt.Println(myBoolean)

	// string
	var myString string = "Hello"
	fmt.Println(myString)

	// int, int8, int 16, int32, int64 (so nguyen) => int8 -> 8 bit => 1 byte
	// int9, int16, int32, int64
	var myInt int = 22
	fmt.Println(myInt)
	// 1. Range
	fmt.Println(math.MinInt)
	fmt.Println(math.MaxInt)

	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt8)

	fmt.Println(math.MinInt16)
	fmt.Println(math.MaxInt16)

	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxInt32)

	fmt.Println(math.MinInt64)
	fmt.Println(math.MaxInt64)

	fmt.Printf("format 'myInt' is of type %T \n", myInt)

	// 2. Bits
	fmt.Println(bits.OnesCount(math.MaxUint))
	fmt.Println(bits.OnesCount8(math.MaxUint8))
	fmt.Println(bits.OnesCount16(math.MaxUint16))
	fmt.Println(bits.OnesCount32(math.MaxUint32))
	fmt.Println(bits.OnesCount64(math.MaxUint64))

	// uint, uint8, uint16, uint32, uint64, uintptr (so nguyen duong)
	var myUnint uint = 20
	fmt.Println(myUnint)

	// byte alias for uint8
	var myByte byte = 'A'
	fmt.Println(myByte)
	fmt.Printf("%T \n", myByte)

	// rune alias for int32
	var myRune rune = 'A'
	fmt.Println(myRune)
	fmt.Printf("%T \n", myRune)

	var myStrig string = "Nhẫn"
	runes := []rune(myStrig)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}
	//represents a Unicode code point

	// float32 float64 (so thuc)
	var myFloat float32 = 45.8
	fmt.Println(myFloat)
	// complex64 complex128 (so phuc) z = a + bi
	var z1 complex64 = 10 + 1i
	fmt.Println(z1)
	var z2 complex64 = 10 + 1i
	fmt.Println(z1 + z2)

	// const
	const pi float64 = 3.1
	// Untyped Constant
	const myFavLanguage = "Python"
	const sunRisesInTheEast = true

	// Multiple declaration
	const country, code = "India", 91

	const (
		employeeId string  = "E101"
		salary     float64 = 50000.0
	)

	// Typed Constant
	const typedInt int = 1234
	const typedStr string = "Hi"

	fmt.Println(myFavLanguage, sunRisesInTheEast, country, code, employeeId, salary, typedInt, typedStr)
	fmt.Println(pi)
}
