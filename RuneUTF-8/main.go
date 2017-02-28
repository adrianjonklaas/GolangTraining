package main

import "fmt"

func main() {
	for i := 1; i <= 350; i++ {
		fmt.Print(i, " - ", string(i), " - ", []byte(string(i)))
		fmt.Printf(" %b \n", i)
	}
	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	foo2 := 'a'
	fmt.Println(foo2)
	fmt.Printf("%T \n", foo2)
}

/*
NOTE:
Some operating systems (Windows) might not print characters where i < 256

If you have this issue, you can use this code:

fmt.Println(i, " - ", string(i), " - ", []int32(string(i)))

UTF-8 is the text coding scheme used by Go.

UTF-8 works with 1 - 4 bytes.

A byte is 8 bits.

[]byte deals with bytes, that is, only 1 byte (8 bits) at a time.

[]int32 allows us to store the value of 4 bytes, that is, 4 bytes * 8 bits per byte = 32 bits.
*/
