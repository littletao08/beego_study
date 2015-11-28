package db
import (
	"github.com/astaxie/beego/utils/pagination"
	"beego_study/utils"
)

type Pagination struct {
	pagination.Paginator
	Data   []interface{}
	hasNext bool
}

func NewPagination(perPageNum int, total int, hasNext bool) *Pagination {
	pagination := new(Pagination)
	pagination.SetNums(total)
	pagination.PerPageNums = perPageNum
	pagination.hasNext = hasNext
	return pagination
}

func ( p *Pagination) HasNext() bool {
	return p.hasNext
}

func (p *Pagination) SetData(container interface{}) {
	p.Data = utils.ToSlice(container)
}





