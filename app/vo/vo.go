package vo

type PagingVo struct {
	Page int `json:"page" form:"page" binding:"omitempty,gt=0"`
	Size int `json:"size" form:"size" binding:"omitempty,gt=0"`
}

type PagingBackVo[T any] struct {
	Data  []T
	Count int
}
