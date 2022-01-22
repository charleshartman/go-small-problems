// 637. Average of Levels in Binary Tree

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) (results []float64) {
	queue := []*TreeNode{root}
	if root == nil {
		return nil
	}

	for len(queue) != 0 {
		level := []int{}
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			currentNode := queue[0]
			queue = queue[1:]
			level = append(level, currentNode.Val)
			left, right := currentNode.Left, currentNode.Right
			if left != nil {
				queue = append(queue, left)
			}
			if right != nil {
				queue = append(queue, right)
			}
		}

		results = append(results, Average(level))
	}

	return
}

func Average(nums []int) float64 {
	var sum float64
	for _, v := range nums {
		sum += float64(v)
	}

	return sum / float64(len(nums))
}
