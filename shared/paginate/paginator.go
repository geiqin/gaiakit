package paginate

import "github.com/geiqin/gaiakit/kitex_gen/base"

type Paginator struct {
	Paged     int64 `json:"paged"`
	Total     int64 `json:"total"`
	PageCount int64 `json:"page_count"`
	PageSize  int64 `json:"page_size"`
	PrevPage  int64 `json:"prev_page"`
	LastPage  int64 `json:"last_page"`
}

func New(paged int64, pageSize ...int64) *Paginator {
	entity := &Paginator{}
	entity.Paged = paged
	if pageSize != nil {
		entity.PageSize = pageSize[0]
	}
	entity.calculate()
	return entity
}

func (a *Paginator) Offset() int {
	offset := (a.Paged - 1) * a.PageSize
	return int(offset)
}

func (a *Paginator) Limit() int {
	return int(a.PageSize)
}

func (a *Paginator) calculate() {
	if a.Paged < 1 {
		a.Paged = 1
	}
	if a.PageSize <= 0 {
		a.PageSize = 20
	}
	a.PageCount = (a.Total + a.PageSize - 1) / a.PageSize
	a.LastPage = a.Paged + 1
	a.PrevPage = a.Paged - 1
	if a.LastPage > a.PageCount {
		a.LastPage = a.PageCount
	}
	if a.PrevPage < 1 {
		a.PrevPage = 1
	}
}

func (a *Paginator) ToPager() *base.Pager {
	a.calculate()
	return &base.Pager{
		Paged:     a.Paged,
		Total:     a.Total,
		PageCount: a.PageCount,
		PageSize:  a.PageSize,
		PrevPage:  a.PrevPage,
		LastPage:  a.LastPage,
	}
}
