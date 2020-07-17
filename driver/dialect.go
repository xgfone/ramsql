package ramsql

import (
	"fmt"
	"strings"
)

// Dialect is used to represent the RAM dialect.
type Dialect struct{}

// Name returns the name of the dialect.
func (d Dialect) Name() string { return "ramsql" }

// Placeholder returns the placeholder of the ith argument.
func (d Dialect) Placeholder(i int) string { return fmt.Sprintf("$%d", i) }

func (d Dialect) quoted(s string) bool   { return strings.IndexByte(s, '"') >= 0 }
func (d Dialect) quote2(s string) string { return fmt.Sprintf(`"%s"`, s) }
func (d Dialect) quote1(s string) string {
	if s == "*" || d.quoted(s) {
		return s
	}

	if i := strings.IndexByte(s, '('); i > -1 {
		_s := s[i+1:]
		if strings.IndexByte(_s, '(') > -1 {
			return s
		}

		i2 := strings.IndexByte(_s, ')')
		if i2 < 0 {
			return s
		}

		return fmt.Sprintf("%s(%s)%s", s[:i], d.quote2(_s[:i2]), _s[i2+1:])
	}

	return d.quote2(s)
}

// Quote returns the quoting string of s.
func (d Dialect) Quote(s string) string {
	if strings.IndexByte(s, ' ') >= 0 {
		return s
	}

	if strings.IndexByte(s, '.') < 0 {
		return d.quote1(s)
	}

	vs := strings.Split(s, ".")
	for i, v := range vs {
		vs[i] = d.quote1(v)
	}

	return strings.Join(vs, ".")
}

// LimitOffset returns the string format of LIMIT and OFFSET.
func (d Dialect) LimitOffset(limit, offset int64) string {
	if limit < 0 {
		panic("sqlx: the limit must be a positive integer")
	}
	if offset == 0 {
		return fmt.Sprintf("LIMIT %d", limit)
	}
	return fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset)
}
