package v1

import "github.com/kamilsk/form-api/pkg/domain"

// PostRequest represents `POST /api/v1/{Schema.ID}` request.
type PostRequest struct {
	ID           domain.ID
	InputData    domain.InputData
	InputContext domain.InputContext
}

// PostResponse represents `POST /api/v1/{Schema.ID}` response.
type PostResponse struct {
	ID     domain.ID
	Error  error
	Schema domain.Schema
}
