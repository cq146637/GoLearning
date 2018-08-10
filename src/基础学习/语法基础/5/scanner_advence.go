package main

import "fmt"

type Ustr string

func (u *Ustr) Scan(state fmt.ScanState, verb rune) (err error) {
	var s []byte
	switch verb {
	case 'S':
		s, err = state.Token(true, func(c rune) bool { return 'A' <= c && c <= 'Z' })
		if err != nil {
			return
		}
	case 's', 'v':
		s, err = state.Token(true, func(c rune) bool { return 'a' <= c && c <= 'z' })
		if err != nil {
			return
		}
	default:
		return fmt.Errorf("无效格式：%c", verb)
	}
	*u = Ustr(s)
	return nil
}

func main() {
	var a, b, c, d, e Ustr
	n, err := fmt.Scanf("%3S%S%3s%2v%x", &a, &b, &c, &d, &e)
	fmt.Println(a, b, c, d, e)
	fmt.Println(n, err)
	// 在终端执行后，输入 ABCDEFGabcdefg 回车
	// 结果：
	// ABC DEFG abc de
	// 4 无效格式：x
}
