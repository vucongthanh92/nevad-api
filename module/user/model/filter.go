package usermodel

type Filter struct {
	OwnerId   int    `json:"owner_id,omitempty" form:"owner_id"`
	Status    []int  `json:"-"`
	SortOrder string `json:"sort_order,omitempty" form:"sort_order"`
	SortField string `json:"sort_field,omitempty" form:"sort_field"`
}
