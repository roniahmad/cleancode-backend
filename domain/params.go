package domain

type (
	Login struct {
		Email    string `form:"email" validate:"required,email"`
		Password string `form:"password" validate:"required"`
	}

	ViewProducts struct {
		Page  int `form:"page" validate:"required"`
		Limit int `form:"limit" validate:"required"`
	}

	CancelOrder struct {
		OrderId int `form:"order_id" validate:"required"`
	}

	CheckoutOrder struct {
		OrderId int `form:"order_id" validate:"required"`
	}

	RemoveOrderItem struct {
		OrderId   int `form:"order_id" validate:"required"`
		ProductId int `form:"product_id" validate:"required"`
	}

	ModifyQuantityOrderItem struct {
		OrderId   int `form:"order_id" validate:"required"`
		ProductId int `form:"product_id" validate:"required"`
		Quantity  int `form:"quantity" validate:"required"`
	}

	BankTrans struct {
		AccNumber string  `form:"acc_number" validate:"required"`
		Amount    float64 `form:"amount" validate:"required"`
	}
)
