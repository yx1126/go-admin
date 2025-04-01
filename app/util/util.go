package util

type TreeNode[T any] interface {
	GetID() int
	GetParentID() *int
	SetChildren(children []TreeNode[T])
}

func ListToTree[T TreeNode[T]](data []T) []T {
	root := make([]T, 0)
	for _, v := range data {
		if v.GetParentID() == nil {
			root = append(root, v)
		} else {
		}
	}
	return root
}
