package dto

type SortBy struct {
	Field string
	Dir   string
}

type FilterBy struct {
	Field    string
	Operator string
	Value    interface{}
}

type SortAndFilterRequest struct {
	Sorts   SortBy
	Filters []FilterBy
	Limit   int
	Offset  int
}

type FilterSql struct {
	SqlOperator  string
	ValueReplace string
}

var Operators = []string{"eq", "neq", "ct", "sw", "ew", "gt", "lt", "gte", "lte"}
