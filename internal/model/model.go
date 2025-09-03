package model

type WebResponse[T any] struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    T             `json:"data,omitempty"`
	Meta    *PageMetadata `json:"meta,omitempty"`
	Errors  string        `json:"errors,omitempty"`
}

type PageResponse[T any] struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    []T          `json:"data,omitempty"`
	Meta    PageMetadata `json:"meta,omitempty"`
	Errors  string       `json:"errors,omitempty"`
}

type PageMetadata struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	LastPage    int64 `json:"last_page"`
	From        int   `json:"from"`
	To          int   `json:"to"`
}

// SwaggerWebResponse is used for Swagger documentation only
type SwaggerWebResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    interface{}   `json:"data,omitempty"`
	Meta    *PageMetadata `json:"meta,omitempty"`
	Errors  string        `json:"errors,omitempty"`
}

// SwaggerPageResponse is used for Swagger documentation only
type SwaggerPageResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data,omitempty"`
	Meta    PageMetadata  `json:"meta,omitempty"`
	Errors  string        `json:"errors,omitempty"`
}
