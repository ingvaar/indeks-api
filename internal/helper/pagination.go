package helper

// PaginateResponse returns a paginated response for the specified data type.
type PaginateResponse[T any] struct {
	Page int `json:"page"`
	Size int `json:"size"`
	TotalPage int `json:"total_page"`
	Data T `json:"data"`
}