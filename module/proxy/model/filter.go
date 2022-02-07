package proxymodel

type Filter struct {
	Status    []int  `json:"-"`
	SortOrder string `json:"sort_order,omitempty" form:"sort_order"`
	SortField string `json:"sort_field,omitempty" form:"sort_field"`
}
