package main

import (
	"fmt"
)

func main() {

	// STRINGS PACKAGE
	// greeting := "Hello there friends!"

	// fmt.Println(strings.Contains(greeting, "friends")) // true, checks if the string contains the substring "friends"
	// fmt.Println(strings.Contains(greeting, "friends!!"))
	// fmt.Println(strings.Contains(greeting, "Friends"))

	// fmt.Println(strings.ToUpper(greeting)) // converts the string to uppercase
	// fmt.Println(strings.Index(greeting, "th"))
	// fmt.Println(strings.Split(greeting, " ")) // splits the string into a slice of strings using the space character as the delimiter

	// fmt.Println(strings.ReplaceAll(greeting, "friends", "ninjas")) // replaces all occurrences of "friends" with "ninjas". This returns a new string, it does not modify the original string.
	// fmt.Println("the original string is still:", greeting)         // the original string is still unchanged

	// SORT PACKAGE
	// ages := []int{20, 30, 10, 40, 50}

	// index2 := sort.SearchInts(ages, 21)
	// fmt.Println("the index of 21 is:", index2) // prints the index of the value 21 in the sorted slice. If searchInts is used on an unsorted slice, result will be incorrect or unexpected.

	// sort.Ints(ages) // sorts the slice of integers in ascending order. Updates the original slice, does not return a new slice.
	// fmt.Println(ages)

	// index3 := sort.SearchInts(ages, 30)        // searches for the index of the value 30 in the sorted slice.
	// fmt.Println("the index of 30 is:", index3) // prints the index of the value 30 in the sorted slice

	// index := sort.SearchInts(ages, 35)        // searches for the index of the value 30 in the sorted slice.
	// fmt.Println("the index of 35 is:", index) // prints the index of the value 30 in the sorted slice

	// names := []string{"Alice", "Bob", "Charlie", "Dave", "Eve"}
	// sort.Strings(names) // sorts the slice of strings in ascending order. Updates the original slice, does not return a new slice.
	// fmt.Println("sorted names:", names)

	// fmt.Println("the index of 'Charlie' is:", sort.SearchStrings(names, "Charlie")) // searches for the index of the value "Charlie" in the sorted slice of strings
	// fmt.Println("the index of 'Frank' is:", sort.SearchStrings(names, "Frank"))     // searches for the index of the value "Frank" in the sorted slice of strings

	// // STANDARD LIBRARY
	// greeting := "hello there friends!"

	// fmt.Println(strings.Contains(greeting, "hello"))
	// fmt.Println(strings.Contains(greeting, "bye"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "ll"))
	// fmt.Println(strings.Split(greeting, " "))

	// // the original value is unchanged
	// fmt.Println("original string value =", greeting)

	// ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	// sort.Ints(ages)
	// fmt.Println(ages) // original slice is changed

	// index := sort.SearchInts(ages, 30)
	// fmt.Println(index)

	// names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	// sort.Strings(names)
	// fmt.Println(names)

	// fmt.Println(sort.SearchStrings(names, "bowser"))

	// // LOOPS
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++ // infinite loop without this
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is:", i)
	// }

	// names := []string{"mario", "luigi", "yoshi", "peach"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, val := range names {
	// 	fmt.Printf("the value at pos %v is %v \n", index, val)
	// 	val = "new string"
	// }

	// for _, val := range names { // If you don't want to use the index, you can use the blank identifier _
	// 	// this will not change the original slice
	// 	// val is a copy of the value in the slice, not a reference to it
	// 	fmt.Print(val, ",")
	// 	val = "new string"
	// }

	// // changing val in a loop does not mutate the original slice
	// fmt.Println(names)

	// BOOLEANS AND CONDITIONALS

	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 40")
	}

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, val := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v \n", index, val)
	}

}
