package usecase

import (
	"context"
	"time"

	"wetees.com/domain"
)

type orderUsecase struct {
	repo           domain.OrderRepository
	conf           *domain.Config
	contextTimeout time.Duration
}

// Checkout Order
func (o *orderUsecase) CheckoutOrder(c context.Context, orderId int, userId int) ([]domain.OrderDetails, int, float64, error) {
	ctx, cancel := context.WithTimeout(c, o.contextTimeout)
	defer cancel()

	orderItems := []domain.OrderDetails{}

	items, err := o.repo.CheckoutOrder(ctx, orderId, userId)
	if err != nil {
		return orderItems, 0, 0, err
	}

	totalItems := 0
	totalPrice := 0.0

	for _, v := range items {
		totalItems += v.Quantity
		totalPrice += (v.Price * float64(v.Quantity))
	}

	return items, totalItems, totalPrice, nil
}

/*
Status Order
NEW (New Order): Customer create a new order
PND (Pending): Customer has begun the checkout process without making the necessary payments for the products.
AWP (Awaiting Payment): Customer may have initiated the payment process but is yet to pay for the product.
PYR (Payment Received): Customer has completed the payment for the order.
ORC (Order Confirmed): Customer has completed the payment and the order has been received and acknowledged by the e-commerce site.
FLD (Failed): Customer could not complete the payment or other verifications required to complete the order.
EXP (Expired): Customer could not make the payment for the products within the stipulated payment window.
AWF (Awaiting Fulfillment): Customer has made the required payments for the price of the products, and the products shall now be shipped.
AWS (Awaiting Shipment): Products bought by the customer are now in a queue ready to be shipped and are waiting to be collected by the shipment service provider.
ONH (On Hold): Stock inventory is reduced by the number of products the customer has requested. However, other steps need to be completed for order fulfillment.
SHP (Shipped): Shipment provider has collected the products and the products are on their way to the customer.
PSP (Partially Shipped): Only a part of the order or some products in the order are shipped.
AWC (Awaiting Pickup): Products have been shipped to either the customer-specified location or the business-specified location and are waiting to be picked up by the customer for delivery.
CPD (Completed): Product has been shipped and delivered, and the payment for the same has been made. The customer, at this point, can receive an invoice regarding the product they bought.
CNC (Canceled): Variety of things. Both the seller and the customer may cancel an order. An order generally shows canceled if the customer fails to make the payment or if the seller has run out of stock of a particular product. Whether or not the customer is entitled to a refund of their money, in this case, depends on the stage of the order and other variables.
DEC (Declined): Seller declares that they cannot ship and fulfill the order.
REF (Refunded): Seller agrees to refund the amount paid by the customer to buy the product.
PRF (Partially Refunded): Seller partially refunds the amount paid by the customer while buying the product.
RFR (Refund Rejected): Seller refuses to process the entire or partial refund of the amount paid by the customer at the time of buying the products.
DPD (Disputed): Customer has raised an issue with the order fulfillment or the refund procedure. Generally, customers raise disputes when e-commerce websites refuse to refund the amount paid by them.
*/

// Create Order
func (o *orderUsecase) CreateOrder(c context.Context, userId int) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(c, o.contextTimeout)
	defer cancel()

	var (
		result map[string]interface{}
	)

	order := domain.Order{
		UserId: userId,
		Status: "NEW",
	}
	id, err := o.repo.CreateOrder(ctx, &order)
	if err != nil {
		return result, err
	}

	result = map[string]interface{}{
		"order_id": id,
	}

	return result, nil
}

// CancelOrder implements domain.OrderUsecase.
func (uc *orderUsecase) CancelOrder(c context.Context, orderId int, userId int) error {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	err := uc.repo.DeleteOrder(ctx, orderId, userId)

	return err
}

func NewOrderUsecase(repo domain.OrderRepository, conf *domain.Config, timeout time.Duration) domain.OrderUsecase {
	return &orderUsecase{
		repo:           repo,
		conf:           conf,
		contextTimeout: timeout,
	}
}
