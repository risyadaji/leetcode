package main

// Ref: https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/

type Solution struct {
	Nodes []int
}

// using in order to get ascending nodes
func (s *Solution) inOrder(root *TreeNode) {
	if root != nil {
		s.inOrder(root.Left)
		s.Nodes = append(s.Nodes, root.Val)
		s.inOrder(root.Right)
	}
}

func kthSmallest(root *TreeNode, k int) int {
	s := Solution{}
	s.inOrder(root)

	return s.Nodes[k-1]
}
