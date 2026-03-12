package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[0:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"

}

func main() {
	f1, s1 := getInitials("hello rusiru")
	fmt.Println(f1, s1)

	f2, s2 := getInitials("hi devinda")
	fmt.Println(f2, s2)

	f3, s3 := getInitials("hey")
	fmt.Println(f3, s3)
}
