package test

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestSlice(t *testing.T) {
	list1 := make([]int, 0, 8)
	arrar := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	list2 := arrar[1:5]

	listappend(list1)
	listappend(list2)
	fmt.Println(list1, len(list1), cap(list1))
	fmt.Println(list2, len(list2), cap(list2))
	fmt.Println(arrar, len(arrar), cap(arrar))
	temp := make([]int, 0)
	fmt.Println("***************", len(temp))
}

func listappend(list []int) {
	h := (*[3]uintptr)(unsafe.Pointer(&list))
	fmt.Println("++++", h[0], h[1], h[2], &h[0], &h[1], &h[2], list)
	list = append(list, 1)
	fmt.Println("++++", h[0], h[1], h[2], &h[0], &h[1], &h[2], list)
}

func TestList(t *testing.T) {
	test1 := make([]int, 0, 8)
	test1 = append(test1, 1)
	h := (*[3]uintptr)(unsafe.Pointer(&test1))
	fmt.Println("----", h[0], h[1], h[2], &h[0], &h[1], &h[2], test1)
	listappend(test1)
	fmt.Println("----", h[0], h[1], h[2], &h[0], &h[1], &h[2], test1)
}
