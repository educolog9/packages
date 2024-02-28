package functions

import (
	"fmt"
	"strconv"

	"github.com/educolog7/packages/enums"
	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Offset int64
	Limit  int64
	Search string
	Sort   string
	Order  enums.SortOrder
}

func ParsePaginationParams(c *gin.Context) (*Pagination, error) {
	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")
	search := c.DefaultQuery("search", "")
	sort := c.DefaultQuery("sort", "created_at")
	orderStr := c.DefaultQuery("order", "asc")

	var offset int64 = 0 // Default offset value
	var err error
	if offsetStr != "" {
		offset, err = strconv.ParseInt(offsetStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid offset format: %w", err)
		}
	}

	var limit int64 = 10 // Default limit value
	if limitStr != "" {
		limit, err = strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid limit format: %w", err)
		}
	}

	var order enums.SortOrder
	if orderStr == "asc" {
		order = enums.Asc
	} else {
		order = enums.Desc
	}

	pagination := &Pagination{
		Offset: offset,
		Limit:  limit,
		Search: search,
		Sort:   sort,
		Order:  order,
	}

	return pagination, nil
}
