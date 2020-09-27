package main

import (
	"crypto/md5"
	"crypto/sha1"   //Go implements several hash functions in various crypto/* packages.
	"crypto/sha256" //Go implements several hash functions in various crypto/* packages.
	"fmt"
)

func main() {
	// 更简单的方法
	fmt.Println("更简单的方法：")
	fmt.Println("md5:", md5.Sum([]byte("hello")))
	fmt.Println("sha1:", sha1.Sum([]byte("hello")))
	fmt.Println("sha256:", sha256.Sum256([]byte("hello")))

	s := "sha1 this string"

	/* The pattern for generating a hash is sha1.New(), sha1.Write(bytes), then sha1.Sum([]byte{}). Here we start with a new hash. */
	h := sha1.New()

	/* Write expects bytes. If you have a string s, use []byte(s) to coerce it to bytes. */
	h.Write([]byte(s))

	//This gets the finalized hash result as a byte slice. The argument to Sum can be used to append to an existing byte slice: it usually isn’t needed.
	bs := h.Sum(nil)

	//SHA1 values are often printed in hex, for example in git commits. Use the %x format verb to convert a hash results to a hex string.
	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	m := md5.New()
	m.Write([]byte(s))
	ms := m.Sum(nil)
	fmt.Println("md5: ", ms)
	fmt.Printf("md5: %x\n", ms)
}
