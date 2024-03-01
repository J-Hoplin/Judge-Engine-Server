package common

// Pagination should only comes with Query
type PaginationQuery struct {
	RequestLimit  int `form:"limit"`  // Limit of user request
	RequestOffset int `form:"offset"` // Offset of user request
	QueryOffset   int // Offset for ORM query
	QueryLimit    int // Limit for ORM query
}

// Private Method
func (pagination *PaginationQuery) validatePaginationOption() {
	if pagination.RequestOffset < 0 {
		pagination.RequestOffset = 0
	}
	if pagination.RequestLimit <= 0 {
		pagination.RequestLimit = 10
	}
}

// Public Method
func (pagination *PaginationQuery) CalculatePaginationOption() {
	// Validate pagination range
	pagination.validatePaginationOption()
	// Calculate pagination limit
	pagination.QueryOffset = (pagination.RequestOffset * pagination.RequestLimit) + 1
	pagination.QueryLimit = pagination.RequestLimit
}
