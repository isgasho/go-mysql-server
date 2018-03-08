package expression

import (
	"gopkg.in/src-d/go-mysql-server.v0/sql"
)

// Not is a node that negates an expression.
type Not struct {
	UnaryExpression
}

// NewNot returns a new Not node.
func NewNot(child sql.Expression) *Not {
	return &Not{UnaryExpression{child}}
}

// Type implements the Expression interface.
func (e Not) Type() sql.Type {
	return sql.Boolean
}

// Eval implements the Expression interface.
func (e Not) Eval(session sql.Session, row sql.Row) (interface{}, error) {
	v, err := e.Child.Eval(session, row)
	if err != nil {
		return nil, err
	}

	if v == nil {
		return nil, nil
	}

	return !v.(bool), nil
}

// Name implements the Expression interface.
func (e Not) Name() string {
	return "Not(" + e.Child.Name() + ")"
}

// TransformUp implements the Expression interface.
func (e *Not) TransformUp(f func(sql.Expression) sql.Expression) sql.Expression {
	return f(NewNot(e.Child.TransformUp(f)))
}
