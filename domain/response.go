package domain

type (
	// Response with message
	MessageResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	// Response with data
	DataResponse struct {
		Success bool        `json:"success"`
		Payload interface{} `json:"payload"`
	}

	// Response with paginated data
	PaginatedDataResponse struct {
		Success    bool        `json:"success"`
		Data       interface{} `json:"data"`
		Total      int         `json:"total"`
		Page       int         `json:"page"`
		TotalPages int         `json:"total_pages"`
		Limit      int         `json:"limit"`
	}

	// Response with checkout data
	CheckoutDataResponse struct {
		Success    bool        `json:"success"`
		Message    string      `json:"message"`
		Items      interface{} `json:"items"`
		TotalItems int         `json:"total_items"`
		TotalPrice int         `json:"total_price"`
		OrderId    int         `json:"order_id"`
	}
)
