package _45_binaryTreeTail

import (
	"fmt"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(postorderTraversal(root))

	root = &TreeNode{
		Val:   3,
		Left:  &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(postorderTraversal(root))
}
