package test

import (
	"fmt"
	"testing"
	"unsafe"
)

func CheckType(args ...interface{}) {
	if true {
		for _, arg := range args {
			switch arg.(type) {
			case int:
				fmt.Println(arg, "is an int value.")
			case string:
				fmt.Println(arg, "is a string value.")
			case int64:
				fmt.Println(arg, "is an int64 value.")
			case uintptr:
				fmt.Println(arg, "is an uintptr value.")
			case *uintptr:
				fmt.Println(arg, "is an *uintptr value.")
			default:
				fmt.Println(arg, "is an unknown type.")
			}
		}
	}
}
func String2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func demo1() {
	s := "123456"
	b := []byte(s)
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := (*[3]uintptr)(unsafe.Pointer(&b))
	fmt.Println(x[0], x[1])
	fmt.Println(h[0], h[1], h[2])
	CheckType(x[0], x[1], h[0], h[1], h[2])
}

func TestOtherDemo(t *testing.T) {
	yinzhengjie := make(map[int]string)
	letter := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for k, v := range letter {
		yinzhengjie[k] = v
	}
	fmt.Printf("字典中的值为：【%v】\n", yinzhengjie) //注意，字典是无序的哟！
	if v, ok := yinzhengjie[1]; ok {
		fmt.Println("存在key=", v)
	} else {
		fmt.Println("没有找到key=", v)
	}
	// str2 := &str
	// fmt.Println(&str, str2)
	// fmt.Println(&str, str)
	// ret := GetString(str)
	// fmt.Println(&ret, ret)
}
