package db
import (
	"beego_study/utils"
	"math"
	"net/url"
	"strconv"
)

const DEFAULT_PER_PAGE = 10

type Mode int

const (
	FULL Mode = 1 + iota
	NEXT_ONLY
)

type Pagination struct {
	Page      int
	PerPage   int
	Total     int
	Data      []interface{}
	hasNext   bool
	pageRange []int
	url       *url.URL
}

func NewPagination(page int, total int, hasNext bool) *Pagination {
	pagination := new(Pagination)

	if page <= 0 {
		page = 1
	}
	pagination.PerPage = DEFAULT_PER_PAGE
	pagination.Page = page
	pagination.Total = total
	pagination.hasNext = hasNext
	return pagination
}

func (p *Pagination) setPerPage(perPage int) {
	p.PerPage = perPage
}

func (p *Pagination) TotalPages() int {
	return (p.Total + p.PerPage - 1) / p.PerPage;
}

func (p *Pagination) NextPage() int {
	if (p.Page < p.TotalPages()) {
		return p.Page + 1;
	}
	return -1;
}

func (p *Pagination) PrevPage() int {
	if p.Page <= 1 {
		return -1
	}else {
		return p.Page - 1
	}
}

func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.PerPage + 1;
}

func ( p *Pagination) HasNext() bool {
	return p.hasNext
}

func (p *Pagination) SetData(container interface{}) {
	p.Data = utils.ToSlice(container)
}


func (p *Pagination) Pages() []int {
	if p.pageRange == nil && p.Total > 0 {
		var pages []int
		pageNums := p.TotalPages()
		page := p.Page
		switch {
		case page >= pageNums - 4 && pageNums > 9:
			start := pageNums - 9 + 1
			pages = make([]int, 9)
			for i := range pages {
				pages[i] = start + i
			}
		case page >= 5 && pageNums > 9:
			start := page - 5 + 1
			pages = make([]int, int(math.Min(9, float64(page + 4 + 1))))
			for i := range pages {
				pages[i] = start + i
			}
		default:
			pages = make([]int, int(math.Min(9, float64(pageNums))))
			for i := range pages {
				pages[i] = i + 1
			}
		}
		p.pageRange = pages
	}
	return p.pageRange
}


func (p *Pagination) PageLink(page int) string {
	values := p.url.Query()
	values.Set("page", strconv.Itoa(page))
	p.url.RawQuery = values.Encode()
	return p.url.String()
}

// Returns URL to the previous page.
func (p *Pagination) PageLinkPrev() (link string) {
	if p.HasPrev() {
		link = p.PageLink(p.Page - 1)
	}
	return
}

// Returns URL to the next page.
func (p *Pagination) PageLinkNext() (link string) {
	if p.HasNext() {
		link = p.PageLink(p.Page + 1)
	}
	return
}

// Returns URL to the first page.
func (p *Pagination) PageLinkFirst() (link string) {
	return p.PageLink(1)
}

// Returns URL to the last page.
func (p *Pagination) PageLinkLast() (link string) {
	return p.PageLink(p.TotalPages())
}

func ( p *Pagination) HasPrev() bool {
	return p.Page > 1
}

func (p *Pagination) IsActive(pagea int) bool {
	return p.Page == pagea
}

func (p *Pagination) SetUrl(url *url.URL) {
	p.url = url
}









