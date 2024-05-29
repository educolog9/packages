package databases

import (
	"fmt"
	"reflect"

	"github.com/educolog9/packages/enums"
	"github.com/educolog9/packages/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConvertPaginationToMongoFilter converts a PaginationConfig into a MongoDB filter and FindOptions.
// It takes a PaginationConfig as input and returns a bson.M filter, *options.FindOptions, and an error.
// The filter is constructed based on the pagination parameters such as offset, limit, sort, search, and filters.
// The FindOptions are set based on the limit and skip values.
// The function supports various filter operators such as equal, not equal, greater than, greater than or equal to,
// less than, less than or equal to, in, not in, like, and not like.
// If an unsupported operator is encountered, an error is returned.
func ConvertPaginationToMongoFilter(config *types.PaginationConfig) (bson.M, *options.FindOptions, error) {
	findOptions := options.Find()

	withLimit := config.WithLimit
	pagination := config.Pagination

	if withLimit {
		findOptions.SetSkip(pagination.GetOffset())
		findOptions.SetLimit(pagination.GetLimit())
	}

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
		value := f.Value
		if reflect.TypeOf(f.Value).Kind() == reflect.String {
			id, err := primitive.ObjectIDFromHex(f.Value.(string))
			if err == nil {
				value = id
			}
		}

		switch f.Operator {
		case enums.Equal:
			filter[f.Field] = bson.M{"$eq": value}
		case enums.NotEqual:
			filter[f.Field] = bson.M{"$ne": value}
		case enums.GreaterThan:
			filter[f.Field] = bson.M{"$gt": value}
		case enums.GreaterThanOrEqual:
			filter[f.Field] = bson.M{"$gte": value}
		case enums.LessThan:
			filter[f.Field] = bson.M{"$lt": value}
		case enums.LessThanOrEqual:
			filter[f.Field] = bson.M{"$lte": value}
		case enums.In:
			switch v := value.(type) {
			case []string:
				filter[f.Field] = bson.M{"$in": v}
			case []interface{}:
				filter[f.Field] = bson.M{"$in": v}
			case string:
				filter[f.Field] = bson.M{"$in": []string{v}}
			default:
				return nil, nil, fmt.Errorf("invalid format for 'in' operator")
			}
		case enums.NotIn:
			switch v := value.(type) {
			case []string:
				filter[f.Field] = bson.M{"$nin": v}
			case []interface{}:
				filter[f.Field] = bson.M{"$nin": v}
			case string:
				filter[f.Field] = bson.M{"$nin": []string{v}}
			default:
				return nil, nil, fmt.Errorf("invalid format for 'not in' operator")
			}
		case enums.Like:
			filter[f.Field] = bson.M{"$regex": value}
		case enums.NotLike:
			filter[f.Field] = bson.M{"$not": bson.M{"$regex": value}}
		default:
			return nil, nil, fmt.Errorf("unsupported operator %s", f.Operator)
		}
	}

	return filter, findOptions, nil
}

// ConvertPaginationToMongoPipeline converts a PaginationConfig into a MongoDB filter and pipeline.
// It takes a PaginationConfig as input and returns a bson.M filter, []bson.M pipeline, and an error.
func ConvertPaginationToMongoPipeline(config *types.PaginationConfig) (bson.M, mongo.Pipeline, error) {
	var pipeline []bson.D

	withLimit := config.WithLimit
	pagination := config.Pagination

	if withLimit {
		pipeline = append(pipeline, bson.D{{Key: "$skip", Value: pagination.GetOffset()}})
		pipeline = append(pipeline, bson.D{{Key: "$limit", Value: pagination.GetLimit()}})
	}

	if pagination.GetSort() != "" {
		sortOrder := 1
		if pagination.GetOrder() == "desc" {
			sortOrder = -1
		}
		pipeline = append(pipeline, bson.D{{Key: "$sort", Value: bson.D{{Key: pagination.GetSort(), Value: sortOrder}}}})
	}

	filter := bson.M{}
	if pagination.GetSearch() != "" && !config.WithAtlasSearch {
		filter = bson.M{"$text": bson.M{"$search": pagination.GetSearch()}}
	}

	for _, f := range pagination.GetFilters() {
		value := f.Value
		if reflect.TypeOf(f.Value).Kind() == reflect.String {
			id, err := primitive.ObjectIDFromHex(f.Value.(string))
			if err == nil {
				value = id
			}
		}

		switch f.Operator {
		case enums.Equal:
			filter[f.Field] = bson.M{"$eq": value}
		case enums.NotEqual:
			filter[f.Field] = bson.M{"$ne": value}
		case enums.GreaterThan:
			filter[f.Field] = bson.M{"$gt": value}
		case enums.GreaterThanOrEqual:
			filter[f.Field] = bson.M{"$gte": value}
		case enums.LessThan:
			filter[f.Field] = bson.M{"$lt": value}
		case enums.LessThanOrEqual:
			filter[f.Field] = bson.M{"$lte": value}
		case enums.In:
			switch v := value.(type) {
			case []string:
				filter[f.Field] = bson.M{"$in": v}
			case []interface{}:
				filter[f.Field] = bson.M{"$in": v}
			case string:
				filter[f.Field] = bson.M{"$in": []string{v}}
			default:
				return nil, nil, fmt.Errorf("invalid format for 'in' operator")
			}
		case enums.NotIn:
			switch v := value.(type) {
			case []string:
				filter[f.Field] = bson.M{"$nin": v}
			case []interface{}:
				filter[f.Field] = bson.M{"$nin": v}
			case string:
				filter[f.Field] = bson.M{"$nin": []string{v}}
			default:
				return nil, nil, fmt.Errorf("invalid format for 'not in' operator")
			}
		case enums.Like:
			filter[f.Field] = bson.M{"$regex": value}
		case enums.NotLike:
			filter[f.Field] = bson.M{"$not": bson.M{"$regex": value}}
		default:
			return nil, nil, fmt.Errorf("unsupported operator %s", f.Operator)
		}
	}

	return filter, mongo.Pipeline(pipeline), nil
}
