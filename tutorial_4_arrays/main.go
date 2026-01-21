package main

import "fmt"

func main() {
	//fixed length
	//same type
	//indexable
	//contiguous in memory

	var intArr [3]int32      //fixed length of same type (array is 3 integers) (0,0,0 by default)
	intArr[1] = 123          //change index 1
	fmt.Println(intArr[0])   //prints index 0
	fmt.Println(intArr[1:3]) //prints index 2-3

	//prints memory address of each index
	//remember each index value is stored continuously in memory (+4 each time)
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	//alternate ways to declare arrays
	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr2)

	intArr3 := [3]int32{1, 2, 3}
	fmt.Println(intArr3)

	intArr4 := [...]int32{1, 2, 3}
	fmt.Println(intArr4)

	//slices are just arrays with additional abilities (essentially wrappers around arrays)
	//most times better to use than arrays

}
