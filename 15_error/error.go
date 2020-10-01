package main

import (
	"errors"
	"fmt"
)

// 当可能有多个错误类型当时候，可以预定义错误类型，然后根据返回当错误类型来判断发生来什么样的错误。
var err42 = errors.New("cannot work with 42")
var err404 = errors.New("cannot find the page")

func f1(arg int) (int, error) {
	if arg == 42 {
		// return -1, errors.New("can't work with 42")
		return -1, err42
	}
	return arg + 3, nil
}

// 一个struct可以来实现Error()
type argError struct {
	arg  int    // index
	prob string // content
}

func (e *argError) Error() string { //该struct实现了Error()
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"} // 因为该struct实现来Error()的接口，所以可以直接返回该实例
	}
	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok { //断言，判断error是不死argError
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
