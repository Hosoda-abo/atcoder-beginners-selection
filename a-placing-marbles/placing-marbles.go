package main

import "fmt"

func main() {
	var s string
	counter := 0
	fmt.Scan(&s)
	if s[0] == '1' {
		counter++
	}
	if s[1] == '1' {
		counter++
	}
	if s[2] == '1' {
		counter++
	}
	fmt.Println(counter)
}
