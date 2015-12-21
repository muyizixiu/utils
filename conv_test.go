// Package utils provides
package utils

import "testing"
import "fmt"

const num = 10

func Test_string(t *testing.T) {
	var a int = num
	println(String(a))
	var b uint = num
	println(String(b))
}
func Test_Int(t *testing.T) {
	fmt.Println(Int("hello world"))
	fmt.Println(Int("1111"))
	fmt.Println(Number("-111", true))
	fmt.Println(Number("111"))
}
