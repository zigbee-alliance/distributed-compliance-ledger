package main

import "fmt"

type s1 struct {
	X int
	Y int
	sq *s2
}

type s2 struct {
	zz bool
}

func main() {
	fmt.Println(s1{
		X: 3,
	})
}
