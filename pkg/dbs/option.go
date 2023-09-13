package dbs

type option struct {
	query    []Query
	order    any
	offset   int
	limit    int
	preloads []string
}

type Query struct {
	Query string
	Args  []any
}

func NewQuery(query string, args ...any) Query {
	return Query{
		Query: query,
		Args:  args,
	}
}

type optionFn func(*option)

type FindOption interface {
	apply(*option)
}
