package main

import (
	"fmt"
	"testing"
)

func WhileDemo(args ...interface{}) {
	if true {
	LOOP:
		for _, arg := range args {
			switch arg.(type) {
			case int:
				fmt.Println(arg, "is an int value.")
				break LOOP
			case string:
				fmt.Println(arg, "is a string value.")
			case int64:
				fmt.Println(arg, "is an int64 value.")
			default:
				fmt.Println(arg, "is an unknown type.")
			}
		}
		fmt.Println("in")
	}
	fmt.Println("end")
}

func TestOtherDemo(t *testing.T) {
	WhileDemo(1)
}
