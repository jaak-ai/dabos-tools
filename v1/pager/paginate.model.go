package pager

import (
	"math"
)

type PaginateModel[T any] struct {
	TotalDocs  int `json:"totalDocs"`
	TotalPages int `json:"totalPages"`

	DocList []*T `json:"docList"`

	Limit int `json:"limit"`
	Page  int `json:"page"`

	NextPage bool `json:"nextPage"`
	PrevPage bool `json:"prevPage"`
}

func (p *PaginateModel[T]) SetLimit(value int) {
	p.Limit = value
}
func (p *PaginateModel[T]) SetPage(value int) {
	p.Page = value
}
func (p *PaginateModel[T]) SetTotalDocs(value int) {
	p.TotalDocs = value

	if p.Limit > 0 {
		p.TotalPages = int(math.Ceil(float64(value) / float64(p.Limit)))
		p.NextPage = p.Page < p.TotalPages
		p.PrevPage = p.Page > 1
	}

}

func (p *PaginateModel[T]) GetSkip() int {
	return p.Limit * (p.Page - 1)
}
