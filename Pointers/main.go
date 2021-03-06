package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {

	var userChoice int16

	fmt.Println("Type in the No. of which program to run:")
	fmt.Println("1", "\t", "Address of a kind")
	fmt.Println("2", "\t", "Convert Meters to Yards")
	fmt.Println("3", "\t", "Exit")

	fmt.Scan(&userChoice)
	if userChoice == 1 {
		address()
	} else if userChoice == 2 {
		convertMetersToYards()
	} else if userChoice == 3 {
		// Do nothing
	} else {
		fmt.Println("Invalid entry.")
	}
}

func address() {

	var a int = 43
	var b *int = &a
	var replacementA int
	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Print("In decimal ")
	fmt.Printf("%d\n", &a)
	fmt.Println("b's memory address - ", &b)
	fmt.Print("In decimal ")
	fmt.Printf("%d\n", &b)
	fmt.Println("The value in b's memory address - ", b)
	fmt.Print("In decimal ")
	fmt.Printf("%d\n", b)
	fmt.Println("The value in the memory address that b points to - ", *b)
	fmt.Print("What int shall we change 'var a' to: ")
	fmt.Scanln(&replacementA)
	*b = replacementA
	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Print("In decimal ")
	fmt.Printf("%d\n", &a)
	fmt.Println("b's memory address - ", &b)
	fmt.Print("In decimal ")
	fmt.Printf("%d\n", &b)
	fmt.Println("The value in b's memory address - ", b)
	fmt.Print("In decimal ")
	fmt.Printf("%d\n", b)
	fmt.Println("The value in the memory address that b points to - ", *b)
	main()

}

func convertMetersToYards() {

	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, "yards.")
	main()

}
