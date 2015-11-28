package db
import (
	"github.com/astaxie/beego/orm"
)

type QuerySeter interface {
	orm.QuerySeter
}

type DB interface {
	orm.Ormer
	From(table string) *Querier
}

type db struct {
	orm.Ormer
}

func NewDB() DB {
	o := orm.NewOrm()
	d := new(db)
	d.Ormer = o
	return d
}

func (d db) From(table string) *Querier {
	query := NewQuery(d, NewGenericSQLBuilder())
	query.From(table)
	return query
}

