package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

func add(num int) {
}

func main() {
	n := node{value: 1}
	n.left = &node{value: 2}
	n.right = &node{value: 3}
	n.left.left = &node{value: 4}
	n.left.right = &node{value: 5}
	n.right.left = &node{value: 6}
	n.right.right = &node{value: 7}

	fmt.Println(n.value)

	fmt.Println(n.left.value)
}
