package main

import "fmt"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (i == 0 && j == 0) || (i == 0 && j == x-1) || (i == y-1 && j == 0) || (i == y-1 && j == x-1) {
				fmt.Print("o")
			} else if i == 0 || i == y-1 {
				fmt.Print("-")
			} else if j == 0 || j == x-1 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

//func main() {
//	QuadA(5, 5) // 输出一个5x5的矩形
//	QuadA(3, 8) // 输出一个3x8的矩形
//	QuadA(0, 0)
//	QuadA(2, 2)
//	QuadA(2, 0)
//	QuadA(2, 1)
//}
