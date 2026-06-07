package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	isRoot := make(map[int]bool)
	nodes := make(map[int]*TreeNode)

	for _, array := range descriptions {
		parent := array[0]
		child := array[1]
		isRoot[child] = false
		_, ok := isRoot[parent]
		if !ok {
			isRoot[parent] = true
		}
		isLeft := array[2]
		v, ok := nodes[parent]
		childv, childok := nodes[child]
		if ok {
			//如果有
			if isLeft == 1 {
				if childok {
					v.Left = childv
				} else {
					v.Left = &TreeNode{Val: child}
					nodes[child] = v.Left
				}
			} else {
				if childok {
					v.Right = childv
				} else {
					v.Right = &TreeNode{Val: child}
					nodes[child] = v.Right
				}
			}
		} else {
			//如果没有
			nodes[parent] = &TreeNode{Val: parent}
			if isLeft == 1 {
				if childok {
					nodes[parent].Left = childv
				} else {
					nodes[parent].Left = &TreeNode{Val: child}
					nodes[child] = nodes[parent].Left
				}
			} else {
				if childok {
					nodes[parent].Right = childv
				} else {
					nodes[parent].Right = &TreeNode{Val: child}
					nodes[child] = nodes[parent].Right
				}
			}
		}
	}
	for key, val := range isRoot {
		if val {
			return nodes[key]
		}
	}
	return nil
}
