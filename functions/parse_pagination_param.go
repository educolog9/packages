package functions

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/educolog7/packages/enums"
	"github.com/educolog7/packages/types"
	"github.com/gin-gonic/gin"
)

func ParsePaginationParams(c *gin.Context) (*types.Pagination, error) {
	paginationEncode := c.Query("p")

	decodedData, err := base64.URLEncoding.DecodeString(paginationEncode)
	if err != nil {
		return nil, fmt.Errorf("invalid pagination format: %w", err)
	}

	var p types.Pagination
	err = json.Unmarshal(decodedData, &p)
	if err != nil {
		return nil, fmt.Errorf("invalid pagination data: %w", err)
	}

	if p.Order == "" {
		p.Order = enums.Asc
	} else {
		if p.Order != enums.Asc && p.Order != enums.Desc {
			return nil, fmt.Errorf("invalid order format")
		}
	}

	return &p, nil
}
