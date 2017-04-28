package dataStructures

// IntBinaryNode is an integer node in a binary tree.
type IntBinaryNode struct {
	Value       int
	Left, Right *IntBinaryNode
}

// IntBinaryTree is binary tree comprised of IntBinaryNode.
type IntBinaryTree struct {
	head *IntBinaryNode
}

// Add created a new node with the given value and adds it to the tree.
func (tree *IntBinaryTree) Add(val int) {
	newNode := new(IntBinaryNode)
	newNode.Value = val

	if tree.head == nil {
		tree.head = newNode
		return
	}

	point := tree.head

	for {
		if newNode.Value < point.Value {
			if point.Left == nil {
				point.Left = newNode
				break
			} else {
				point = point.Left
			}
		} else {
			if point.Right == nil {
				point.Right = newNode
				break
			} else {
				point = point.Right
			}
		}
	}
}

// VisitAscending recursively traverses the binary tree and applies the visit function to each node.
func (tree *IntBinaryTree) VisitAscending(visit func(node *IntBinaryNode)) {
	visitAscending_rec(tree.head, visit)
}

func visitAscending_rec(node *IntBinaryNode, visit func(node *IntBinaryNode)) {
	if node == nil {
		return
	}

	visitAscending_rec(node.Left, visit)
	visit(node)
	visitAscending_rec(node.Right, visit)
}
