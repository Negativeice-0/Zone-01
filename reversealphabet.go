package main

import "fmt"

func ReverseAlphabet() {
	for i := 'z'; i >= 'a'; i-- {
		fmt.Printf("%c", i)
	}
	fmt.Println()
}

//func main() {
//	ReverseAlphabet()
//}
