package main

import (
	"fmt"
	"testing"
)

func listappend(list []int) {
	list = append(list, 1)
	fmt.Println(list, len(list), cap(list))
}

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
