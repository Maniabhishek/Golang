/*
Text formatting
fmt package provides terminal printing and string formatting
Provides functions
 Printf custom format
 Print simple print
 Println simple print with a newline
F and S variants of the above functions:
 F prints to a data stream : Fprintf , Fprint, Fprintln
 S prints to a new string : SPrintf , Sprint , Sprintln

*/

/*
Printf uses verbs to describe how something should print
Verb      Description
%v        default
%t 		  true or false
%c        character
%X        Hex
%U        Unicode format
%e        Scientific notation
*/
/*
Escape sequence
\\        backslash
\'        single quote
\"        doubole quote
\n        new line
\u or \U  unicode 2 byte & 4 byte
\x        Raw bytes (as hex digits)
*/

package main

import "fmt"

func surrounded(message string, left rune, right rune) string {
	return fmt.Sprintf("%c%v%c", left, message, right)
}

func main() {
	fmt.Println("hello word ")
	fmt.Printf("%v%v \n", "hello", "world")

	fmt.Println(surrounded("hello", '(', ')'))
}
