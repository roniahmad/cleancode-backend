package vars

import "errors"

var (
	ErrChangePasswordFailed  = newError(ChangePasswordFailed)
	ErrCreateOrderFailed     = newError(CreateOrderFailed)
	ErrCreateOrderItemFailed = newError(CreateOrderItemFailed)
	ErrCancelOrderFailed     = newError(CancelOrderFailed)
	ErrCheckoutOrderFailed   = newError(CheckoutOrderFailed)
	ErrChartIsEmpty          = newError(ChartIsEmpty)
	ErrDepositFailed         = newError(DepositFailed)
	ErrExpiredToken          = newError(TokenExpired)
	ErrInvalidCredentials    = newError(InvalidCredentials)
	ErrInvalidAuthHeader     = newError(InvalidAuthHeader)
	ErrInvalidToken          = newError(TokenInvalid)
	ErrLoginFailed           = newError(LoginFailed)
	ErrNotEnoughBalance      = newError(NotEnoughBalance)
	ErrNotFound              = newError(NotFound)
	ErrPasswordNotMatch      = newError(PasswordNotMatch)
	ErrRegisterUserFailed    = newError(RegisterUserFailed)
	ErrRemoveOrderItemFailed = newError(RemoveOrderItemFailed)
	ErrUnAuthorized          = newError(UnAuthorized)
	ErrUserAlreadyExists     = newError(UserAlreadyExists)
	ErrUserWithEmailNotFound = newError(EmailNotFound)
	ErrUpdateQtyFailed       = newError(UpdateQtyFailed)
	ErrWithdrawsFailed       = newError(WithdrawsFailed)
)

const (
	ActivateEmail          = "activate email"
	BadRequest             = "bad Request"
	Bearer                 = "Bearer"
	ChangePasswordFailed   = "change password failed"
	ChangePasswordSucceed  = "change password succeed"
	ChartIsEmpty           = "chart is empty"
	CancelOrderFailed      = "cancel order failed"
	CancelOrderSucceed     = "cancel order succeed"
	CheckoutOrderFailed    = "checkout order failed"
	CheckoutOrderSucceed   = "checkout order succeed"
	CreateOrderFailed      = "create order failed"
	CreateOrderSucceed     = "create order succeed"
	CreateOrderItemFailed  = "create order item failed"
	CreateOrderItemSucceed = "create order item succeed"
	DepositFailed          = "deposit failed"
	DepositSucceed         = "deposit succeed"
	EmailNotFound          = "user not found with the given email"
	InvalidAuthHeader      = "Invalid authorization header format"
	InvalidCredentials     = "invalid credentials"
	LoginFailed            = "login failed"
	ModeDev                = "dev"
	ModeProd               = "prod"
	NotEnoughBalance       = "not enough balance"
	NotFound               = "not found"
	PasswordNotMatch       = "password not match"
	Register               = "register"
	RegisterUserFailed     = "register user failed"
	RegisterUserSucceed    = "register user succeed"
	RemoveOrderItemFailed  = "remove order item failed"
	RemoveOrderItemSucceed = "remove order item succeed"
	TokenExpired           = "token expired"
	TokenInvalid           = "token invalid"
	StatusAWP              = "awaiting for payment"
	UnAuthorized           = "not authorized"
	UserAlreadyExists      = "user already exists"
	UpdateQtyFailed        = "update quantity failed"
	UpdateQtySucceed       = "update quantity succeed"
	WithdrawsFailed        = "withdraws failed"
	WithdrawsSucceed       = "withdraws succeed"
)

func newError(message string) error {
	return errors.New(message)
}
