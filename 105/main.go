package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//输出: [3,9,20,null,null,15,7]
// preorder[0]=root
// preorder[1]=inorder[0] 最左子树 index 回父亲
// inorder[1]=root
// preorder[2]=right1=20(右子树的根)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	idxMap := make(map[int]int, len(inorder))
	for i, v := range inorder {
		idxMap[v] = i
	}
	// preorder 中左右
	// inorder 左中右
	//preorder=[1,2,4,5,3]
	//inorder=[4,2,5,1,3]
	// preL=0,preL=4,inL=0,inR=4
	// root=preorder[0]=1
	// index=3
	// leftSize=index-inL=3
	// 左子树
	// 左子树的前序区间 [2 4 5] preL=1,preR=preL+leftSize-1(根)=3 [1,3]
	// 左子树的中序区间 [4 2 5] inL=0, inR=inL+leftSize-1=2 [0,2]
	// 右子树
	// 右子树的前序区间 [3] preL=preL+leftSize==1+3=4,preR=4 [4,4]
	// 右子树的中序区间 [3] inL=inL+leftSize+1=0+3+1=4,inR=inR [4,4]
	// 左子树根 2
	// index=1
	// leftSize=index-inL=1
	// 左左子树的前序区间 [4] preL=2,preR=2+1-1=2 [2,2]
	// 左左子树的中序区间 [4] inL=0,inR=inL+leftSize-1=0 [0,0]
	// 左右子树的前序区间 [5] preL=preL+leftSize=2+1=3,preR=3 [3,3]
	// 左右子树的中序区间 [5] inL=preL+leftSize+1=2+1+-1=2,inR=inR=2 [2,2]
	var build func(preL, preR, inL, inR int) *TreeNode
	build = func(preL, preR, inL, inR int) *TreeNode {
		// 边界：区间为空返回 nil
		if preL >= len(preorder) || preR >= len(preorder) || inL >= len(inorder) || inR >= len(inorder) || preL > preR || inL > inR {
			return nil
		}
		root := &TreeNode{
			Val: preorder[preL],
		}
		leftStart := preL + 1
		index := idxMap[root.Val]
		leftSize := index - inL
		root.Left = build(leftStart, leftStart+leftSize-1, inL, inL+leftSize-1)
		root.Right = build(leftStart+leftSize, preR, inL+leftSize+1, inR)
		// root = preorder[preL]
		// idx = 在中序里找 root.Val（先线性找，过了再优化成哈希表）
		// leftSize = idx - inL
		// root.Left  = build(左子树的前序区间, 左子树的中序区间)
		// root.Right = build(右子树的前序区间, 右子树的中序区间)
		return root // ← 挂载+返回，一步都不能少
	}
	return build(0, len(preorder)-1, 0, len(inorder)-1)
}
