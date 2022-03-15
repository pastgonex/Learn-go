package simplify_unit_testing

type Result struct {
	Count int
}

func (r Result) Int() int {
	return r.Count
}

type Rows []struct{}

type Stmt interface {
	Close() error
	NumInput() int
	Exec(stmt string, args ...string) (Result, error)
	Query(args []string) (Rows, error)
}
