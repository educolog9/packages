package types

import "github.com/educolog7/packages/enums"

type Filter struct {
	Field    string
	Operator enums.Operator
	Value    interface{}
}

type PaginationConfig struct {
	Pagination *Pagination
	WithLimit  bool
}

func NewPaginationConfig(pagination *Pagination) *PaginationConfig {
	return &PaginationConfig{
		Pagination: pagination,
		WithLimit:  true, // default value
	}
}

type Pagination struct {
	Offset  int64
	Limit   int64
	Search  string
	Sort    string
	Order   enums.SortOrder
	Next    string
	Prev    string
	Filters []Filter
}

func (p *Pagination) GetOffset() int64 {
	return p.Offset
}

func (p *Pagination) GetLimit() int64 {
	if p.Limit == 0 {
		return 10
	}
	return p.Limit
}

func (p *Pagination) GetSearch() string {
	return p.Search
}

func (p *Pagination) GetSort() string {
	return p.Sort
}

func (p *Pagination) GetOrder() enums.SortOrder {
	if p.Order == "" {
		return enums.SortOrder("asc")
	}
	return p.Order
}

func (p *Pagination) GetNext() string {
	return p.Next
}

func (p *Pagination) GetPrev() string {
	return p.Prev
}

func (p *Pagination) GetFilters() []Filter {
	return p.Filters
}
