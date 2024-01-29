package main

import "fmt"

func main() {
var num int

fmt.Print("Enter a number: ")
fmt.Scanln(&num)

for i = '0'; i > num; i-- {
if num%2 == 0 {
fmt.Println("The entered number is even:", num)
} else {
fmt.Println("The entered number is not even.")
}
}
}