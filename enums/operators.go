package enums

type Operator string

// Equal represents the equality operator.
const (
	Equal              Operator = "eq"
	NotEqual           Operator = "ne"
	GreaterThan        Operator = "gt"
	GreaterThanOrEqual Operator = "gte"
	LessThan           Operator = "lt"
	LessThanOrEqual    Operator = "lte"
	Like               Operator = "like"
	NotLike            Operator = "notLike"
	In                 Operator = "in"
	NotIn              Operator = "notIn"
)

// Example of usage in a filter url:

// /api/v1/users?filters=[{"field":"name","operator":"eq","value":"John"}]
// /api/v1/users?filters=[{"field":"age","operator":"gte","value":"18"}]
// /api/v1/users?filters=[{"field":"age","operator":"in","value":"18,19,20"}]
// /api/v1/users?filters=[{"field":"name","operator":"like","value":"John"}]
// /api/v1/users?filters=[{"field":"name","operator":"notLike","value":"John"}]
// /api/v1/users?filters=[{"field":"age","operator":"lt","value":"18"}]
// /api/v1/users?filters=[{"field":"age","operator":"lte","value":"18"}]
// /api/v1/users?filters=[{"field":"age","operator":"ne","value":"18"}]
// /api/v1/users?filters=[{"field":"age","operator":"notIn","value":"18,19,20"}]
// /api/v1/users?filters=[{"field":"age","operator":"gt","value":"18"}]
