package mem

import (
	"fmt"

	"gopkg.in/src-d/go-mysql-server.v0/sql"
)

// Table represents an in-memory database table.
type Table struct {
	name   string
	schema sql.Schema
	data   []sql.Row
}

// NewTable creates a new Table with the given name and schema.
func NewTable(name string, schema sql.Schema) *Table {
	return &Table{
		name:   name,
		schema: schema,
	}
}

// Resolved implements the Resolvable interface.
func (Table) Resolved() bool {
	return true
}

// Name returns the table name.
func (t *Table) Name() string {
	return t.name
}

// Schema implements the Node interface.
func (t *Table) Schema() sql.Schema {
	return t.schema
}

// Children implements the Node interface.
func (t *Table) Children() []sql.Node {
	return []sql.Node{}
}

// RowIter implements the Node interface.
func (t *Table) RowIter(session sql.Session) (sql.RowIter, error) {
	return sql.RowsToRowIter(t.data...), nil
}

// TransformUp implements the Transformer interface.
func (t *Table) TransformUp(f func(sql.Node) sql.Node) sql.Node {
	return f(t)
}

// TransformExpressionsUp implements the Transformer interface.
func (t *Table) TransformExpressionsUp(f func(sql.Expression) sql.Expression) sql.Node {
	return t
}

// Insert a new row into the table.
func (t *Table) Insert(row sql.Row) error {
	if len(row) != len(t.schema) {
		return fmt.Errorf("insert expected %d values, got %d", len(t.schema), len(row))
	}

	for idx, value := range row {
		c := t.schema[idx]
		if !c.Check(value) {
			return sql.ErrInvalidType.New(value)
		}
	}

	t.data = append(t.data, row.Copy())
	return nil
}
