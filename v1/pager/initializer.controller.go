package pager

func New[T any]() *PaginateModel[T] {
	return &PaginateModel[T]{
		TotalDocs:  0,
		TotalPages: 0,
		DocList:    make([]*T, 0),

		Limit: 10,
		Page:  1,

		NextPage: false,
		PrevPage: false,
	}
}
