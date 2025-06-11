package main

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
	// age := 20
	// name := "Mario"

	// fmt.Println("Hello, Ninjas!")
	// fmt.Print("Hello, ")
	// fmt.Print("World! \n")              // Print does not add a new line at the end
	// fmt.Println("Hello again, Ninjas!") // Println adds a new line at the end
	// fmt.Print("Hello again, everyone! \n")

	// //Println with variables
	// fmt.Println("Hello, Ninjas! My name is", name, "and I am", age, "years old.") // I don't need to add spaces between the variables and the strings, Println adds them automatically.

	// // Printf (formatted strings) %_ = format specifier; does not add a new line at the end
	// // %s = string, %d = integer, %f = float, %t = boolean
	// // %v = any value, will automatically convert the value to a string
	// // %T = type of the value
	// // %q = quoted string
	// // and many more, see https://pkg.go.dev/fmt#hdr-Printing
	// fmt.Printf("Hello, Ninjas! My name is %v and I am %v years old.\n", name, age) // %v is a placeholder for any value. It will automatically convert the value to a string.
	// fmt.Printf("Hello, Ninjas! My name is %q and I am %q years old.\n", name, age) // %q is a placeholder for a quoted string. It will automatically convert the value to a string and add quotes around it. Does not work with integers, so it will convert the integer to a string first.
	// fmt.Printf("age is of type %T \n", age)                                        // %T is a placeholder for the type of the value. It will automatically convert the value to a string and print the type of the value.
	// fmt.Printf("you scored %f points! \n", 25.555)
	// fmt.Printf("you scored %0.1f points! \n", 25.55)
	// fmt.Printf("you scored %0.2f points! \n", 25.555) // %.2f will print the float with 2 decimal places. It will round the number to 2 decimal places.

	// // Sprintf (save formatted string to a variable)
	// var message string = fmt.Sprintf("Hello, Ninjas! My name is %v and I am %v years old.", name, age) // Sprintf returns a formatted string, it does not print it to the console.
	// fmt.Println("the saved string is:", message)                                                       // Print the formatted string

	// ARRAYS AND SLICES
	// arrays are fixed size, slices are dynamic size. Slices are arrays under the hood, but they are more flexible and easier to work with. Slices can grow and shrink in size, while arrays cannot.
	// Will probably use slices most of the time, but arrays are still useful in some cases.
	// arrays
	// var ages[3] int = [3]int{20, 30, 40} // array of integers with a fixed size of 3
	// var ages = [3]int{20, 30, 40} // slice of integers with a dynamic size. Slices are more flexible than arrays.

	// names := [4]string{"yoshi", "mario", "peach", "bowser"} // array of strings with a fixed size of 4
	// names[1] = "luigi"                                      // reassigning a value in the array
	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	// // slices
	// var scores []int = []int{100, 90, 80} // slice of integers with a dynamic size
	// fmt.Println(scores, len(scores))
	// scores = append(scores, 70) // append a value to the slice. Slices can grow in size. Arrays can't. Remember it needs to be assigned to a variable, does not automatically append.
	// fmt.Println(scores, len(scores))
	// scores[2] = 85 // reassigning a value in the slice
	// fmt.Println(scores, len(scores))
	// scores = append(scores, 60, 50, 40) // append multiple values to the slice
	// fmt.Println(scores, len(scores))

	// //slice ranges. Picking a range of values from a slice to create a new slice.
	// newScores := scores[1:4]               // create a new slice with values from index 1 to index 3 (not including index 4)
	// fmt.Println(newScores, len(newScores)) // prints the new slice and its length
	// newScores = scores[:3]                 // create a new slice with values from index 0 to index 2 (not including index 3)
	// fmt.Println(newScores, len(newScores)) // prints the new slice and its length
	// newScores = scores[2:]                 // create a new slice with values from index 2 to the end of the slice
	// fmt.Println(newScores, len(newScores)) // prints the new slice and its length
	// newScores = scores[:]                  // create a new slice with all values from the original slice
	// fmt.Println(newScores, len(newScores)) // prints the new slice and its length
	// newScores = scores[2:2]
	// fmt.Println(newScores, len(newScores)) // prints the new slice and its length, which is an empty slice because the start and end index are the same
	// // newScores = scores[2:1] // this will cause a runtime error because the start index is greater than the end index. Slices cannot be created with a start index greater than the end index.
	// newScores = scores[2:8]                   // this will not cause a runtime error, but it will create a new slice with values from index 2 to the end of the slice, even if the end index is greater than the length of the slice.
	// fmt.Println(newScores, len(newScores))    // prints the new slice and its length, which is an empty slice because the end index is greater than the length of the slice
	// newScores = append(newScores, 30, 20, 10) // append values to the new slice
	// fmt.Println(newScores, len(newScores))    // prints the new slice and its length after appending values
}
