package pagination

import (
	"net/http"
	"strconv"
)

const (
	PageSize     = "page_size"
	PageNumber   = "page_number"
	SortField    = "sort_field"
	SortType     = "sort_type"
	DefaultLimit = 30
)

// PagingAndSortParams
type PagingAndSortParams struct {
	PageSize   int
	PageNumber int
	SortField  string
	SortType   string // DESC/ASC
}

// GetPagingAndSortParamsFromRequest will get the pagination and sort params from the request
func GetPagingAndSortParamsFromRequest(r *http.Request) (PagingAndSortParams, error) {
	params := PagingAndSortParams{
		PageSize: DefaultLimit,
	}
	// get url params
	if pageSize := r.URL.Query().Get(PageSize); len(pageSize) > 0 {
		num, err := strconv.Atoi(pageSize)
		if err != nil {
			return PagingAndSortParams{}, err
		}
		params.PageSize = num
	}
	if pageNumber := r.URL.Query().Get(PageNumber); len(pageNumber) > 0 {
		num, err := strconv.Atoi(pageNumber)
		if err != nil {
			return PagingAndSortParams{}, err
		}
		params.PageNumber = num
	}
	if sortField := r.URL.Query().Get(SortField); len(sortField) > 0 {
		params.SortField = sortField
	}
	if sortType := r.URL.Query().Get(SortType); len(sortType) > 0 {
		params.SortType = sortType
	}
	return params, nil
}
