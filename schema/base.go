package schema

// HTTPList HTTP响应列表数据
type HTTPList struct {
	List       interface{}     `json:"list"`
	Pagination *HTTPPagination `json:"pagination,omitempty"`
}

// HTTPPagination HTTP分页数据
type HTTPPagination struct {
	Total     int64 `json:"total"`
	PageIndex int64 `json:"pageIndex"`
	PageSize  int64 `json:"pageSize"`
}
