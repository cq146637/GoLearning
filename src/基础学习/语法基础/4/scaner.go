package main

import "fmt"

// 对于 Scan 而言，回车视为空白
func main() {
	a, b, c := "", 0, false
	fmt.Scan(&a, &b, &c)
	fmt.Println(a, b, c)
	// 在终端执行后，输入 abc 1 回车 true 回车
	// 结果 abc 1 true
}

// 对于 Scanln 而言，回车结束扫描
func main2() {
	a, b, c := "", 0, false
	fmt.Scanln(&a, &b, &c)
	fmt.Println(a, b, c)
	// 在终端执行后，输入 abc 1 true 回车
	// 结果 abc 1 true
}

// 格式字符串可以指定宽度
func main3() {
	a, b, c := "", 0, false
	fmt.Scanf("%4s%d%t", &a, &b, &c)
	fmt.Println(a, b, c)
	// 在终端执行后，输入 1234567true 回车
	// 结果 1234 567 true
}
