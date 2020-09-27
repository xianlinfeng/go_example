package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("%v\n", p) //print an instance of our point struct.
	//OutPut:	{1 2}

	fmt.Printf("%+v\n", p) //If the value is a struct, the %+v variant will include the struct’s field names.
	//OutPut:	{x:1 y:2}

	// fmt.Printf("%#v\n", p)	/The %#v variant prints a Go syntax representation of the value, i.e. the source code snippet that would produce that value.
	//OutPut:	main.point{x:1, y:2}

	fmt.Printf("%T\n", p) //To print the type of a value, use %T.
	//OutPut:	main.point

	fmt.Printf("%t\n", true) //Formatting booleans is straight-forward.
	//OutPut:	true

	fmt.Printf("%d\n", 123) //	 Use %d for standard, base-10 formatting.
	//OutPut:	123

	fmt.Printf("%b\n", 14) // a binary representation.
	//OutPut:	1110

	fmt.Printf("%c\n", 33) //prints the character corresponding to the given integer.
	//OutPut:	!

	fmt.Printf("%x\n", 456) //%x provides hex encoding.
	//OutPut:	1c8

	fmt.Printf("%f\n", 78.9) //There are also several formatting options for floats. For basic decimal formatting use %f.
	//OutPut:	78.900000

	fmt.Printf("%e\n", 123400000.0) //%e and %E format the float in (slightly different versions of) scientific notation.
	fmt.Printf("%E\n", 123400000.0)
	//OutPut:	1.234000e+08
	// OutPut:	1.234000E+08

	fmt.Printf("%s\n", "\"string\"") //For basic string printing use %s.
	//OutPut:	"string"

	fmt.Printf("%q\n", "\"string\"") //To double-quote strings as in Go source, use %q.
	//OutPut:	"\"string\""

	fmt.Printf("%x\n", "hex this") //As with integers seen earlier, %x renders the string in base-16, with two output characters per byte of input.
	//OutPut:	6865782074686973

	fmt.Printf("%p\n", &p) //To print a representation of a pointer, use %p.
	//OutPut:	0xc000016070

	fmt.Printf("|%6d|%6d|\n", 12, 345) // format a integer output
	//OutPut:	|    12|   345|

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) //You can also specify the width of printed floats, though
	//OutPut:	|  1.20|  3.45|

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) //靠左对齐 To left-justify, use the - flag.
	//OutPut:	|1.20  |3.45  |

	fmt.Printf("|%6s|%6s|\n", "foo", "b") // 对string进行对齐
	//OutPut:	|   foo|     b|

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") // 靠左对齐
	//OutPut:	|foo   |b     |

	s := fmt.Sprintf("a %s", "string")
	//So far we’ve seen Printf, which prints the formatted string to os.Stdout.
	//Sprintf formats and returns a string without printing it anywhere.
	fmt.Println(s)
	//OutPut:	a string

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
	//You can format+print to io.Writers other than os.Stdout using Fprintf.
	//OutPut:	an error
}
