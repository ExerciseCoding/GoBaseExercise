package leetcode

import (
	"fmt"
	"math"
	"testing"
)

/**

 */
type TreeNode struct {
	val int
	right *TreeNode
	left *TreeNode
}

func maxDepthCount(root *TreeNode)int{
	if root == nil{
		return 0
	}
	fmt.Println("遍历的节点:",root.val)
	ret := int(math.Max(float64(maxDepthCount(root.left)),float64(maxDepthCount(root.right))))+1
	fmt.Println(ret)
	return ret

}

func TestMaxDepthCount(t *testing.T){
	node1 := &TreeNode{val:3}
	node2 := &TreeNode{val:9}
	node3 := &TreeNode{val:20}
	node4 := &TreeNode{val:15}
	node5 := &TreeNode{val:7}

	node1.left = node2
	node1.right = node3
	node3.left = node4
	node3.right = node5

	ret := maxDepthCount(node1)
	fmt.Println("二叉树的深度:",ret)
}
