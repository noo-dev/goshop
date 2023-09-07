package paging

import "math"

const (
	DEFAULT_PAGE_SIZE int64 = 20
)

type Pagination struct {
	CurrentPage int64 `json:"current_page"`
	Total       int64 `json:"total"`
	TotalPage   int64 `json:"total_page"`
	Limit       int64 `json:"limit"`
	Skip        int64 `json:"skip"`
}

func New(page int64, pageSize int64, total int64) *Pagination {
	var pagination Pagination
	limit := DEFAULT_PAGE_SIZE
	if pageSize > 0 && pageSize <= limit {
		pagination.Limit = pageSize
	} else {
		pagination.Limit = limit
	}

	pagination.TotalPage = int64(math.Ceil(float64(total) / float64(pagination.Limit)))
	pagination.Total = total
	if page < 1 || pagination.TotalPage == 0 {
		page = 1
	}
	pagination.CurrentPage = page
	pagination.Skip = (page - 1) * pagination.Limit
	return &pagination
}
