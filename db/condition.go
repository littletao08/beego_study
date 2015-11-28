package db
import (
"errors"
	"fmt"
)

type ConditionType int

const (
	SEGMENT ConditionType = 1 + iota
	EQ
	NOT_EQ
	IN
	NOT_IN
	LIKE
	NOT_LIKE
	BETWEEN
	NOT_BETWEEN
	LESS
	LE
	GREAT
	GE
	NULL
	NOT_NULL
)

var conditionTypes = [...]string{
	"",
	"=",
	"!=",
	"",
	"",
	"like",
	"not like",
	"",
	"",
	"<",
	"<=",
	">",
	">=",
	"",
	"",
}

func (c ConditionType) String() string {
	return conditionTypes[c - 1]
}

type Condition struct {
	Type ConditionType
	Column    string
	Params    []interface{}
}

func NewCondition(condition ConditionType, column string, param interface{}) Condition {
	var c Condition
	c.Type = condition
	c.Column = column
	c.Params = toParams(param)
	return c
}

func toParams(param interface{}) []interface{} {
	if nil == param {
		return []interface{}{}
	}

	params, ok := param.([]interface{})
	fmt.Println("**********ok***********",ok)
	if ok {
		return params
	}else {
		return []interface{}{param}
	}
}

func (c Condition) ToSQL(sqlBuilder SQLBuilder) (string, error) {
	if len(c.Type.String()) >0 {
       return sqlBuilder.EscapeColumn(c.Column) + " " + c.Type.String() + " ?",nil
    }
	switch c.Type {
	case SEGMENT:
		return "(" + c.Column + ")",errors.New("segment analysis error")
	case IN:
	case NOT_IN:
		var sql = sqlBuilder.EscapeColumn(c.Column)
		if c.Type == NOT_IN {
			sql += " not "
		}
		sql += "in ("
		for i, _ := range c.Params {
			if (i > 0) {
				sql += ", "
			}
			sql += "?"
		}
		sql += ")"
		return sql,errors.New("not in analysis error")
	case BETWEEN:
	case NOT_BETWEEN:
		var sql = sqlBuilder.EscapeColumn(c.Column)
		if c.Type == NOT_BETWEEN {
			sql += " not "
		}
		sql += " between ? and ? "
		return sql,errors.New("not between analysis error")
	case NULL:
		return sqlBuilder.EscapeColumn(c.Column) + " is null",errors.New("null analysis error");
	case NOT_NULL:
		return sqlBuilder.EscapeColumn(c.Column) + " is not null",errors.New("not null analysis error");
	}

	return "", errors.New("sql analysis error")
}
