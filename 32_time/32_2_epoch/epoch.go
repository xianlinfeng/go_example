/*A common requirement in programs is getting the number of seconds, milliseconds, or nanoseconds since the
Unix epoch. Here’s how to do it in Go. */
package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	secs := now.Unix() // unix.epoch后经过了多少second
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs/2, 0)) // 将seconds转换为unix.epoch
	fmt.Println(time.Unix(0, nanos))  // 将nanoseconds转换为unix.epoch

}
