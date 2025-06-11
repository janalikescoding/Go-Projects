package main

import "fmt"

func main() {
	// FIRST PROGRAM
	//fmt.Println("Hello, Ninjas!")

	// STRINGS
	// var nameOne string = "mario"
	// var nameTwo = "luigi"
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// // reassigning variable
	// nameOne = "peach"
	// nameThree = "bowser"

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameFour := "yoshi" // short variable declaration. Can be used only for new variables. Can be used only inside functions.
	// fmt.Println(nameFour)

	// ints (positive and negative integers, no decimals)
	// var ageOne int = 20
	// var ageTwo = 30
	// ageThree := 40 // short variable declaration
	// fmt.Println(ageOne, ageTwo, ageThree)

	// var numOne int8 = 25  // 8 bits. When specified, it can only hold values from -128 to 127.
	// var numTwo int8 = -128
	// var numThree uint8 = 255 // unsigned int, can only hold values from 0 to 255

	// floats (decimal numbers). Need to specify memory (directly or indirectly).
	// var priceOne float32 = 19.99 // 32 bits
	// var priceTwo float64 = 296546545464654654654654654.99 // 64 bits, more precise
	// scoreThree := 99.99 // short variable declaration, defaults to float64

	//PRINT
	age := 20
	name := "Mario"

	fmt.Println("Hello, Ninjas!")
	fmt.Print("Hello, ")
	fmt.Print("World! \n")              // Print does not add a new line at the end
	fmt.Println("Hello again, Ninjas!") // Println adds a new line at the end
	fmt.Print("Hello again, everyone! \n")

	//Println with variables
	fmt.Println("Hello, Ninjas! My name is", name, "and I am", age, "years old.") // I don't need to add spaces between the variables and the strings, Println adds them automatically.

	// Printf (formatted strings) %_ = format specifier; does not add a new line at the end
	// %s = string, %d = integer, %f = float, %t = boolean
	// %v = any value, will automatically convert the value to a string
	// %T = type of the value
	// %q = quoted string
	// and many more, see https://pkg.go.dev/fmt#hdr-Printing
	fmt.Printf("Hello, Ninjas! My name is %v and I am %v years old.\n", name, age) // %v is a placeholder for any value. It will automatically convert the value to a string.
	fmt.Printf("Hello, Ninjas! My name is %q and I am %q years old.\n", name, age) // %q is a placeholder for a quoted string. It will automatically convert the value to a string and add quotes around it. Does not work with integers, so it will convert the integer to a string first.
	fmt.Printf("age is of type %T \n", age)                                        // %T is a placeholder for the type of the value. It will automatically convert the value to a string and print the type of the value.
	fmt.Printf("you scored %f points! \n", 25.555)
	fmt.Printf("you scored %0.1f points! \n", 25.55)
	fmt.Printf("you scored %0.2f points! \n", 25.555) // %.2f will print the float with 2 decimal places. It will round the number to 2 decimal places.

	// Sprintf (save formatted string to a variable)
	var message string = fmt.Sprintf("Hello, Ninjas! My name is %v and I am %v years old.", name, age) // Sprintf returns a formatted string, it does not print it to the console.
	fmt.Println("the saved string is:", message)                                                       // Print the formatted string
}
