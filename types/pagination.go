package types

import "github.com/educolog7/packages/enums"

type Pagination struct {
	Offset int64
	Limit  int64
	Search string
	Sort   string
	Order  enums.SortOrder
}
