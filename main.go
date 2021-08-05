package main

import "fmt"

var name = 18

func main() {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println(test())
	fmt.Println(name)
}
func test() int {
	defer add()
	name = name + 1
	return name
}
func add() {
	name = name + 1
}

