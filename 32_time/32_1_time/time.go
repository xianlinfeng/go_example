package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now() //当前时间
	p(now)
	p(now.Year())
	p(now.Month())
	p(now.Day())
	p(now.Weekday())

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC) // 指定时间
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	p(then.Before(now)) // 时间比较
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then) // 时间差
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff)) // **之后
	p(then.Add(-diff))
}
