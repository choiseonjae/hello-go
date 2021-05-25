package main

import (
	"math"
	"strconv"
)

func main()  {
	// overflow error test
	/*
	var integer32 int32 = 2147483648
	println(integer32)

	var uint32 uint32 = -10
	println(uint32)
	 */

	var float32 float32 = 2147483647
	var float64 float64 = 9223372036854775807
	println(float32, float64)

	// 두 형식의 제한 확인 가능
	println(math.MaxFloat32, math.MaxFloat64)

	var firstName string = "John"
	lastName := "Doe"
	println(firstName, lastName)
	lastName = "Tom"
	println(firstName, lastName)

	// type convert
	var i int = 1
 	println(i)
	var s string = strconv.Itoa(i)
	println(s)
}
