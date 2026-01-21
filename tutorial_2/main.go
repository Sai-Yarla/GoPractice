package main

import "fmt"

func main() {
	// default value of int is 0, also for all integer types and float types and runes
	// default value of bool is false
	// default value of string is ""
	var intNum int
	fmt.Println(intNum)

	var floatNum float64 = 1234567.89
	fmt.Println(floatNum) // prints: 1234567.89000

	//no mixed type arithmetic operations

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = float32(intNum32) + floatNum32
	fmt.Println(result) // prints: 20.2

	//integer division results in a integer rounded down and reamainder with modulus operator
	var a int = 7
	var b int = 3
	var divResult int = a / b
	var modResult int = a % b
	fmt.Println(divResult) // prints: 2
	fmt.Println(modResult) // prints: 1

	var myString string = "test"
	fmt.Println(myString)

	//we can drop the type declaration and let go infer the type if a value is assigned
	var myString2 = "test2"
	fmt.Println(myString2)

	// we can drop var and use := to declare and assign a variable inside functions
	myVar := "test3" // only works inside functions
	fmt.Println(myVar)

	//prints 4
	fmt.Println((len("test"))) //prints the number of BYTES, NOT number of characters

	//prints 1
	fmt.Println(len("A"))

	//len is fine for basic characters but may not work for special characters

	var myBool bool = true
	fmt.Println(myBool)

	//value of const can not be changed
	//must declare type and initialize it every time
	const myConst string = "const value"
	fmt.Println(myConst)

}
