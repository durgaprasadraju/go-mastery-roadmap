package solutions

import "strings"

// BuildSelect builds a parameterized SELECT query.
func BuildSelect(table string, columns []string, where string) string {
	q := "SELECT " + strings.Join(columns, ", ") + " FROM " + table
	if where != "" {
		q += " WHERE " + where
	}
	return q
}

// QuoteIdentifier escapes a SQL identifier (simplified).
func QuoteIdentifier(name string) string {
	return "\"" + strings.ReplaceAll(name, "\"", "\"\"") + "\""
}