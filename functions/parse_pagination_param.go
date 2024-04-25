package functions

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/educolog9/packages/enums"
	"github.com/educolog9/packages/types"
	"github.com/gin-gonic/gin"
)

func ParsePaginationParams(c *gin.Context) (*types.Pagination, error) {
	paginationEncode := c.Query("p")

	var p types.Pagination

	if paginationEncode == "" {
		p.Limit = 10
		p.Order = enums.Asc
	} else {
		decodedData, err := base64.URLEncoding.DecodeString(paginationEncode)
		if err != nil {
			return nil, fmt.Errorf("invalid pagination format: %w", err)
		}

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
	}

	return &p, nil
}
