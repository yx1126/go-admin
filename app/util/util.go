package util

type TreeNode[T any] interface {
	GetID() int
	GetParentID() int
	SetChildren(children []T) T
}

func ListToTree[T TreeNode[T]](data []T, parentId int) []T {
	tree := make([]T, 0)
	if len(data) <= 0 {
		return tree
	}
	// 处理顶级节点 parentId不为0的情况
	if parentId == -1 {
		parentId = data[0].GetParentID()
		for _, v := range data[1:] {
			if parentId < v.GetParentID() {
				parentId = v.GetParentID()
			}
		}
	}
	for _, menu := range data {
		if menu.GetParentID() == parentId {
			tree = append(tree, menu.SetChildren(ListToTree(data, menu.GetID())))
		}
	}
	return tree
}
