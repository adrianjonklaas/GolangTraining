package main

import "fmt"

type contact struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	var userChoice int16

	fmt.Println("1", "\t", "Loop: Print Odd Numbers from 1-50")
	fmt.Println("2", "\t", "Loop: Print Sum of Numbers Divisible by 3 and 5 from 1-1000")
	fmt.Println("3", "\t", "Switch: Favourite Friends")
	fmt.Println("4", "\t", "Switch: Switch on Type")
	fmt.Println("5", "\t", "Exit")
	fmt.Print("Type in the No. of which program to run: ")

	fmt.Scan(&userChoice)
	if userChoice == 1 {
		loopOddNumbers()
	} else if userChoice == 2 {
		loopSumDivThreeFive()
	} else if userChoice == 3 {
		switchFriends()
	} else if userChoice == 4 {
		switchType()
	} else if userChoice == 5 {
		// Do nothing
	} else {
		fmt.Println("Invalid entry.")
	}
}

func loopOddNumbers() {

	//This loop demonstrates a for loop with no condition; it also demonstrates the use of Continue and break
	/* Idiomatically, var i int is preferred to i := 0; since i is incremented immediately within the for loop
	it is the equivalent to starting with i = 1 */

	var i int

	for {
		i++
		if i%2 == 0 {
			continue
		}
		if i >= 50 {
			break
		} else {
			fmt.Println(i)
		}
	}
	main()
}

func loopSumDivThreeFive() {

	var counter int

	for i := 1; i < 1000; i++ {
		if i%3 == 0 && i%5 == 0 {
			counter += i
		}
	}
	fmt.Println(counter)
	main()
}

func switchFriends() {

	var bestFriend string
	fmt.Print("Enter your best friend's name: ")
	fmt.Scanln(&bestFriend)

	switch bestFriend {
	case "Dami":
		fmt.Println("Wassup Dami")
	case "Jobby":
		fmt.Println("Wassup Jobby")
	case "Jimmy":
		fmt.Println("Wassup Jimmy")
	case "Vikas":
		fmt.Println("Wassup Vikas")
		fallthrough
	default:
		fmt.Println("Sorry, everyone else! There can only be one!")
	}
	main()
}

func switchType() {

	var friend = contact{"Jim,", "Zach", 40}

	SwitchOnType(50)
	SwitchOnType("Adrian")
	SwitchOnType(friend)
	SwitchOnType(friend.firstName)
	SwitchOnType(friend.lastName)
	main()
}

func SwitchOnType(x interface{}) {

	// SwitchOnType works with interfaces

	switch x.(type) { // this is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")

	}
}
