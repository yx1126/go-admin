package vo

type PagingVo struct {
	Page int `json:"page" form:"page" binding:"omitempty,gte=1"`
	Size int `json:"size" form:"size" binding:"omitempty,gte=1"`
}

type PagingBackVo[T any] struct {
	Data  []T
	Count int
}
