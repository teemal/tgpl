package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	t := ""
	//Go doesn't allow unused variables
	//since we are getting an index and index value from range os.Args
	//we can just pass the index into a blank identifier(_) since it won't be used
	for i, arg := range os.Args[1:] {
		t = " index: " + strconv.Itoa(i)
		fmt.Println(arg + t)
	}

	//less constly version of echo
	fmt.Println(strings.Join(os.Args[1:], " "))
}
