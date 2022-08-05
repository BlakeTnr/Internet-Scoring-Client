package api

import "strings"

type Query struct {
	Queries map[string]string
}

func ParseRawQuery(rawQuery string) Query {
	splitQueries := strings.Split(rawQuery, "&")

	queries := make(map[string]string)
	for _, query := range splitQueries {
		splitQuery := strings.Split(query, "=")
		key := splitQuery[0]
		value := splitQuery[1]
		queries[key] = value
	}

	return Query{queries}
}

func (query *Query) GetQuery(key string) string {
	return query.Queries[key]
}
