package main

import (
	"fmt"
	"strconv"
	//"time"
)

//Lazy Evaluation
func Interge() <-chan int {
	yield := make(chan int, 12)
	count := 0
	go func() {
		for {
			yield <- count
			fmt.Println("gorount:" + strconv.Itoa(count))
			count++
		}
	}()
	return yield
}

var resume <-chan int

func getInetge() int {
	return <-resume
}

func main() {
	resume = Interge()
	//	time.Sleep(time.Second * 5)
	fmt.Println(getInetge())
	//time.Sleep(time.Second * 5)
	fmt.Println(getInetge())
	fmt.Println(getInetge())
	fmt.Println(getInetge())
	fmt.Println(getInetge())
	fmt.Println(getInetge())
	fmt.Println(getInetge())
}
