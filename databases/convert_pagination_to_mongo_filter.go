package databases

import (
	"fmt"

	"github.com/educolog7/packages/enums"
	"github.com/educolog7/packages/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConvertPaginationToMongoFilter converts a types.Pagination object into a MongoDB filter and find options.
// It takes the pagination object as input and returns the corresponding filter, find options, and an error (if any).
// The filter is a bson.M object representing the filter conditions based on the pagination parameters.
// The find options are options.FindOptions object that specify skip, limit, and sort options for the MongoDB query.
// If any unsupported operator is encountered in the pagination filters, an error is returned.
func ConvertPaginationToMongoFilter(pagination types.Pagination) (bson.M, *options.FindOptions, error) {
	findOptions := options.Find()
	findOptions.SetSkip(pagination.GetOffset())
	findOptions.SetLimit(pagination.GetLimit())

	if pagination.GetSort() != "" {
		sortOrder := 1
		if pagination.GetOrder() == "desc" {
			sortOrder = -1
		}
		findOptions.SetSort(bson.D{{Key: pagination.GetSort(), Value: sortOrder}})
	}

	filter := bson.M{}
	if pagination.GetSearch() != "" {
		filter = bson.M{"$text": bson.M{"$search": pagination.GetSearch()}}
	}

	for _, f := range pagination.GetFilters() {
		switch f.Operator {
		case enums.Equal:
			filter[f.Field] = bson.M{"$eq": f.Value}
		case enums.NotEqual:
			filter[f.Field] = bson.M{"$ne": f.Value}
		case enums.GreaterThan:
			filter[f.Field] = bson.M{"$gt": f.Value}
		case enums.GreaterThanOrEqual:
			filter[f.Field] = bson.M{"$gte": f.Value}
		case enums.LessThan:
			filter[f.Field] = bson.M{"$lt": f.Value}
		case enums.LessThanOrEqual:
			filter[f.Field] = bson.M{"$lte": f.Value}
		case enums.In:
			values, ok := f.Value.([]interface{})
			if !ok {
				return nil, nil, fmt.Errorf("invalid format for 'in' operator")
			}
			filter[f.Field] = bson.M{"$in": values}
		case enums.NotIn:
			values, ok := f.Value.([]interface{})
			if !ok {
				return nil, nil, fmt.Errorf("invalid format for 'nin' operator")
			}
			filter[f.Field] = bson.M{"$nin": values}
		case enums.Like:
			filter[f.Field] = bson.M{"$regex": f.Value}
		case enums.NotLike:
			filter[f.Field] = bson.M{"$not": bson.M{"$regex": f.Value}}
		default:
			return nil, nil, fmt.Errorf("unsupported operator %s", f.Operator)
		}
	}

	return filter, findOptions, nil
}
