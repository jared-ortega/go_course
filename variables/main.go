package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var myIntVar int
	myIntVar = -12
	fmt.Println("myIntVar: ", myIntVar)

	var myUintVar uint
	myUintVar = 12
	fmt.Println("myUintVar: ", myUintVar)

	var myStringVar string
	myStringVar = "My string var"
	fmt.Println("myStringVar", myStringVar)

	var myBoolVar bool
	myBoolVar = false
	fmt.Println("myBoolVar: ", myBoolVar)

	fmt.Println("myStringVar dir: ", &myStringVar)

	myIntVar2 := 12
	fmt.Println("myIntVar2", myIntVar2)

	const myFirstConst = "a12"
	fmt.Println("My const is: ", myFirstConst)

	const myIntConst int = 12
	fmt.Println("My myIntConst is: ", myIntConst)

	const myStringConst string = "aaaa"
	fmt.Println("My myStringConst is: ", myStringConst)

	fmt.Println()
	var my8BitsIntVar int8
	fmt.Printf("Int default value: %d\n", my8BitsIntVar)
	fmt.Printf("type: %T, bytes: %d, bites %d\n", my8BitsIntVar, unsafe.Sizeof(my8BitsIntVar), unsafe.Sizeof(my8BitsIntVar)*8)

	var my16BitsIntVar int16
	fmt.Printf("Int16 default value: %d\n", my16BitsIntVar)
	fmt.Printf("16type: %T, bytes: %d, bites %d\n", my16BitsIntVar, unsafe.Sizeof(my16BitsIntVar), unsafe.Sizeof(my16BitsIntVar)*8)

	var my64BitsIntVar int64
	fmt.Printf("Int64 default value: %d\n", my64BitsIntVar)
	fmt.Printf("64type: %T, bytes: %d, bites %d\n", my64BitsIntVar, unsafe.Sizeof(my64BitsIntVar), unsafe.Sizeof(my64BitsIntVar)*8)

}
