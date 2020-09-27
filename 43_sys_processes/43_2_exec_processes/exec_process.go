/*
In the previous example we looked at spawning external processes. We do this when we need an external
process accessible to a running Go process. Sometimes we just want to completely replace the current
Go process with another (perhaps non-Go) one. To do this we’ll use Go’s implementation of the classic
exec function.
我们只想用另一个（也许是非Go）进程完全替换当前的Go进程
*/
package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// go需要知道该二进制文件的绝对路径，所以需要用LookPath来找到该路径。
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// define arguments
	args := []string{"ls", "-a", "-l", "-h"}

	// 我们同样也需要environment variables来运行进程
	env := os.Environ()

	//  If this call is successful, the execution of our process will end here
	//		and be replaced by the /bin/ls -a -l -h process.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
