package query

import "github.com/kamilsk/form-api/pkg/domain"

// CreateTemplate TODO
type CreateTemplate struct {
	ID         *domain.ID
	Title      string
	Definition string
}

// ReadTemplate TODO
type ReadTemplate struct {
	ID domain.ID
}

// UpdateTemplate TODO
type UpdateTemplate struct {
	ID         domain.ID
	Title      string
	Definition string
}

// DeleteTemplate TODO
type DeleteTemplate struct {
	ID          domain.ID
	Permanently bool
}