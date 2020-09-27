package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	/*
		直接写入文件
	*/
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	/*
		打开一个文件用来写入
	*/
	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	/* You can Write byte slices as you’d expect. */
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	/* A WriteString is also available. */
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	/* Issue a Sync to flush writes to stable storage.
	发出同步以将写入刷新到稳定的存储。 */
	f.Sync()

	/*
		带有缓存器到writer
	*/
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush() // 用flush来保证所有buffer中到操作，都被运行

}
