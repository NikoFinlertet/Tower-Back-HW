package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}
type Tree struct{ root *Node }

// add вставляет новый узел в бинарное дерево поиска
func (t *Tree) add(num int) {
	p := t.root

	//!!! t.root - error!!!
	if t.root == nil {
		t.root = &Node{value: num}
		return
	}

	for { //Идем циклом
		if num <= p.value { //Если меньше, то идем слева
			if p.left == nil {
				p.left = &Node{value: num}
				return
			}
			p = p.left
		} else {
			if p.right == nil {
				p.right = &Node{value: num}
				return
			}
			p = p.right
		}
	}
}

func (t *Tree) deleteTree() {
	t.root = nil
}

func (t *Tree) delete(num int) {
	p := t.root
	pred := p

	if t.root == nil {
		return
	}

	for {
		//Find
		if num < p.value {
			if p.left != nil {
				pred = p
				p = p.left
			}
		} else if num > p.value {
			if p.right != nil {
				pred = p
				p = p.right
			}
		} else {
			//If find --> DELETE
			if p.left == nil && p.right == nil {
				if pred.left == p {
					pred.left = nil
				} else {
					pred.right = nil
				}
				return

			} else if p.left == nil {
				if pred.left == p {
					pred.left = p.right
				} else {
					pred.right = p.right
				}

			} else if p.right == nil {
				if pred.left == p {
					pred.left = p.left
				} else {
					pred.right = p.left
				}

			} else {
				min := minimumNode(p.right, pred)
				p.value = min.value
				return
			}
		}
	}
}

func minimumNode(this *Node, last *Node) *Node {
	p := this
	pred := last

	for p.left != nil {
		pred = p
		p = p.left
	}

	if pred.left == p {
		pred.left = nil
	}

	return p
}

func (t *Tree) isExist(num int) bool {
	p := t.root

	if t.root == nil {
		return false
	}

	for {
		if num == p.value {
			return true
		}

		if num < p.value {
			if p.left != nil {
				p = p.left
			}
		} else if num < p.value {
			if p.right != nil {
				p = p.right
			}
		} else {
			return false
		}
	}

}

func (t *Tree) print() {
	if t.root != nil {
		t.root.print()
	} else {
		fmt.Println("Tree not exist")
	}
	fmt.Println("")
}

func (n *Node) print() {
	if n != nil {
		n.left.print()
		fmt.Print(n.value, " ")
		n.right.print()
	}
}

func main() {
	A := Tree{}
	B := Tree{}

	A.add(50)
	A.add(30)
	A.add(20)
	A.add(40)
	A.add(70)
	A.add(60)

	B.add(1)
	B.add(599)
	B.add(199)
	B.add(2545)

	fmt.Println("A:")
	A.print()
	fmt.Println("B:")
	B.print()

	fmt.Println("exist '50' is", A.isExist(50))
	A.delete(50)
	fmt.Println("exist '50' is", A.isExist(50))
	A.print()

	B.deleteTree()
	B.print()

}
