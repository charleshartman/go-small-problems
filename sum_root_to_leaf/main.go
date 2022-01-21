/* 129. Sum Root to Leaf Numbers

You are given the root of a binary tree containing digits from 0 to 9 only.

Each root-to-leaf path in the tree represents a number.

For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.
Return the total sum of all root-to-leaf numbers. Test cases are generated so
that the answer will fit in a 32-bit integer.

A leaf node is a node with no children.

Example 1:
Input: root = [1,2,3]
Output: 25
Explanation:
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.
Therefore, sum = 12 + 13 = 25.

Example 2:
Input: root = [4,9,0,5,1]
Output: 1026
Explanation:
The root-to-leaf path 4->9->5 represents the number 495.
The root-to-leaf path 4->9->1 represents the number 491.
The root-to-leaf path 4->0 represents the number 40.
Therefore, sum = 495 + 491 + 40 = 1026.

Constraints:
The number of nodes in the tree is in the range [1, 1000].
0 <= Node.val <= 9
The depth of the tree will not exceed 10.
*/

/*
Sum root of a node is:
  - Current value of the node * 10 + sum root of the left node
  - Current value of the node * 10 + sum root of the right node
  - Add these two up

If node is a leaf node (left and right are nil):
  - Return current value
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//[4,9,0,5,1]
/*
1st Iteration:
  acc = 0 * 10 + 4
  leftSide = helper(node 9, 40)
            -> acc = 4 * 10 + 9 = 49
            -> helper(node 5, 49)
               -> acc = 49 * 10 + 5

*/

func sumNumbers(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, acc int) int {
	if root.Left == nil && root.Right == nil {
		return acc*10 + root.Val
	}

	var leftSide int
	var rightSide int

	acc = (acc * 10) + root.Val

	if root.Left != nil {
		leftSide = helper(root.Left, acc)
	}

	if root.Right != nil {
		rightSide = helper(root.Right, acc)
	}

	return leftSide + rightSide
}
