package main

import "strconv"

func main() {
	println(sum("1", "2"))
	println(calc("1", "2"))
	name := "David"
	notPoint(name)
	println(name)
	point(&name)
	println(name)
}

func sum(num1 string, num2 string) int {
	int1, _ := strconv.Atoi(num1)
	int2, _ := strconv.Atoi(num2)
	return int1 + int2
}

func calc(num1 string, num2 string) (int, int) {
	int1, _ := strconv.Atoi(num1)
	int2, _ := strconv.Atoi(num2)
	return int1 + int2, int1 * int2
}

func notPoint(name string) {
	name = "John"
}

func point(name *string)  {
	*name = "John"
}
