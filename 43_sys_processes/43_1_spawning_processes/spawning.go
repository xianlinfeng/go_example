/*创建外部进程 */
package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	//执行外部命令“date”
	dateCmd := exec.Command("date")

	// 获取输出结果
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut)) // 打印输出结果

	// 执行另一条命令 “grep helle“。
	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()   // define stdin管道
	grepOut, _ := grepCmd.StdoutPipe() // define stdout管道
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep")) // 从管道里面输入 该string
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut) // 输出命令结果
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 也可以在exec.Command()中输入完整命令， 如果想让一个string作为argument，则需要使用 “-c” 标签，然后加相应的string
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
