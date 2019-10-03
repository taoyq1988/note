package tree

import (
	"errors"
)

// value -1 means default value
type BST struct {
	Left *BST
	Right *BST
	Value int
}

func NewBST() *BST {
	root := &BST{
		Value: -1,
	}
	return root
}

// 1. 如果节点 c 为空，则节点 c 的父节点将作为节点 n 的父节点。如果节点 n 的值小于该父节点的值，则节点 n 将作为该父节点的左孩子；否则节点 n 将作为该父节点的右孩子。
// 2. 比较节点 c 与节点 n 的值。
// 3. 如果节点 c 的值与节点 n 的值相等，则说明用户在试图插入一个重复的节点。解决办法可以是直接丢弃节点 n，或者可以抛出异常。
// 4. 如果节点 n 的值小于节点 c 的值，则说明节点 n 一定是在节点 c 的左子树中。则将父节点设置为节点 c，并将节点 c 设置为节点 c 的左孩子，然后返回至第 1 步。
// 5. 如果节点 n 的值大于节点 c 的值，则说明节点 n 一定是在节点 c 的右子树中。则将父节点设置为节点 c，并将节点 c 设置为节点 c 的右孩子，然后返回至第 1 步。
func(root *BST) insert(value int) error {
	if value < 0 {
		return errors.New("illegle value")
	}
	if root.Value == -1 {
		root.Value = value
	}

	return nil
}
