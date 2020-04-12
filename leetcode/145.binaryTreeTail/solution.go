package _45_binaryTreeTail

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	stack := NewStack()
	stack.Push(root)
	for !stack.Empty() {
		if node := stack.Top(); node != nil {
			stack.Pop()
			result = append(result, node.Val)
			if node.Right != nil || node.Left != nil {
				stack.Push(node.Left)
				stack.Push(node.Right)
			}
		} else {
			stack.Pop()
		}
	}
	lenR := len(result)
	for i, j := 0, lenR-1 ; i<j ; i,j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

/**
Stack struct
 */
type Stack struct {
	data  []*TreeNode
	index int
}

func NewStack() *Stack {
	return &Stack{
		data:  make([]*TreeNode, 0),
		index: -1,
	}
}

func (s *Stack) Empty() bool {
	return s.index == -1
}

func (s *Stack) Top() *TreeNode {
	if s.index == -1 {
		return nil
	}
	return s.data[s.index]
}

func (s *Stack) Pop() (d *TreeNode) {
	if s.index == -1 {
		return nil
	}
	d = s.data[s.index]
	s.data = s.data[:s.index]
	s.index--
	return
}

func (s *Stack) Push(d *TreeNode) {
	s.index++
	s.data = append(s.data, d)
}
