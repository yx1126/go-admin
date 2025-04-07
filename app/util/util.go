package util

type TreeNode[T any] interface {
	GetID() int
	GetParentID() int
	SetChildren(children []T) T
}

func ListToTree[T TreeNode[T]](data []T, parentId int) []T {
	tree := make([]T, 0)
	for _, menu := range data {
		if menu.GetParentID() == parentId {
			tree = append(tree, menu.SetChildren(ListToTree(data, menu.GetID())))
		}
	}
	return tree
}
