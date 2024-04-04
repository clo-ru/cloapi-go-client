package clo

import "fmt"

type FilteringField struct {
	FieldName string
	Condition string
	Value     string
}

func AddFilterToRequest(r RequestInt, ff FilteringField) {
	switch ff.Condition {
	case "gt", "gte", "lt", "lte", "range", "in":
		condString := fmt.Sprintf("%s__%s", ff.FieldName, ff.Condition)
		r.WithQueryParams(QueryParam{condString: {ff.Value}})
	}
}
