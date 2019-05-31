package test

import (
	"fmt"
	"testing"
	"time"
)

//http://docscn.studygolang.com/pkg/time/
const TimeFormat = "2006-01-02 15:04:05"

func TestTime(t *testing.T) {
	time.AfterFunc(time.Millisecond,
		func() {
			fmt.Println("243")
			if true {
				fmt.Println("23333")
			}
		},
	)
}

//2006-01-02 15:04:05 07 --- 年 月 日 时 分 秒 时区
func TestFmort(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format(time.RFC3339Nano))
	fmt.Println(now.Format("2006-Jan-2"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))
}

func TestConstruct(t *testing.T) {
	t1, _ := time.Parse("2006-01-02 15:04:05", "2013-08-11 11:18:46")
	fmt.Println(t1)

	t2, _ := time.Parse("2006-01-02:15", "2019-05-24:13")
	fmt.Println(t2)

	t3 := time.Unix(1362984425, 0)
	fmt.Println(t3)
}
