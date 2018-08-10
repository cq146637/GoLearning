package main

import "fmt"

func main() {
	fmt.Print("a", "b", 1, 2, 3, "c", "d", "\n")
	fmt.Println("a", "b", 1, 2, 3, "c", "d")
	fmt.Printf("ab %d %d %d cd\n", 1, 2, 3)
	// ab1 2 3cd
	// a b 1 2 3 c d
	// ab 1 2 3 cd

	if err := percent(30, 70, 90, 160); err != nil {
		fmt.Println(err)
	}
	// 30%
	// 70%
	// 90%
	// 数值 160 超出范围（100）
}

func percent(i ...int) error {
	for _, n := range i {
		if n > 100 {
			return fmt.Errorf("数值 %d 超出范围（100）", n)
		}
		fmt.Print(n, "%\n")
	}
	return nil
}
