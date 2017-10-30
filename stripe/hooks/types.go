package hooks

// EventType is the type of event for the hook.
type EventType string

// Event types available.
const (
	AccountUpdated                    EventType = "account.updated"                      // Occurs whenever an account status or property has changed.
	AccountApplicationDeauthorized              = "account.application.deauthorized"     // Occurs whenever a user deauthorizes an application. Sent to the related application only.
	AccountExternalAccountCreated               = "account.external_account.created"     // Occurs whenever an external account is created.
	AccountExternalAccountDeleted               = "account.external_account.deleted"     // Occurs whenever an external account is deleted.
	AccountExternalAccountUpdated               = "account.external_account.updated"     // Occurs whenever an external account is updated.
	ApplicationFeeCreated                       = "application_fee.created"              // Occurs whenever an application fee is created on a charge.
	ApplicationFeeRefunded                      = "application_fee.refunded"             // Occurs whenever an application fee is refunded, whether from refunding a charge or from refunding the application fee directly, including partial refunds.
	ApplicationFeeRefundUpdated                 = "application_fee.refund.updated"       // Occurs whenever an application fee refund is updated.
	BalanceAvailable                            = "balance.available"                    // Occurs whenever your Stripe balance has been updated (e.g. when a charge collected is available to be paid out). By default, Stripe will automatically transfer any funds in your balance to your bank account on a daily basis.
	BitcoinReceiverCreated                      = "bitcoin.receiver.created"             // Occurs whenever a receiver has been created.
	BitcoinReceiverFilled                       = "bitcoin.receiver.filled"              // Occurs whenever a receiver is filled (that is, when it has received enough bitcoin to process a payment of the same amount).
	BitcoinReceiverUpdated                      = "bitcoin.receiver.updated"             // Occurs whenever a receiver is updated.
	BitcoinReceiverTransactionCreated           = "bitcoin.receiver.transaction.created" // Occurs whenever bitcoin is pushed to a receiver.
	ChargeCaptured                              = "charge.captured"                      // Occurs whenever a previously uncaptured charge is captured.
	ChargeFailed                                = "charge.failed"                        // Occurs whenever a failed charge attempt occurs.
	ChargeRefunded                              = "charge.refunded"                      // Occurs whenever a charge is refunded, including partial refunds.
	ChargeSucceeded                             = "charge.succeeded"                     // Occurs whenever a new charge is created and is successful.
	ChargeUpdated                               = "charge.updated"                       // Occurs whenever a charge description or metadata is updated.
	ChargeDisputeClosed                         = "charge.dispute.closed"                // Occurs when the dispute is closed and the dispute status changes to charge_refunded, lost, warning_closed, or won.
	ChargeDisputeCreated                        = "charge.dispute.created"               // Occurs whenever a customer disputes a charge with their bank (chargeback).
	ChargeDisputeFundsReinstated                = "charge.dispute.funds_reinstated"      // Occurs when funds are reinstated to your account after a dispute is won.
	ChargeDisputeFundsWithdrawn                 = "charge.dispute.funds_withdrawn"       // Occurs when funds are removed from your account due to a dispute.
	ChargeDisputeUpdated                        = "charge.dispute.updated"               // Occurs when the dispute is updated (usually with evidence).
	CouponCreated                               = "coupon.created"                       // Occurs whenever a coupon is created.
	CouponDeleted                               = "coupon.deleted"                       // Occurs whenever a coupon is deleted.
	CouponUpdated                               = "coupon.updated"                       // Occurs whenever a coupon is updated.
	CustomerCreated                             = "customer.created"                     // Occurs whenever a new customer is created.
	CustomerDeleted                             = "customer.deleted"                     // Occurs whenever a customer is deleted.
	CustomerUpdated                             = "customer.updated"                     // Occurs whenever any property of a customer changes.
	CustomerDiscountCreated                     = "customer.discount.created"            // Occurs whenever a coupon is attached to a customer.
	CustomerDiscountDeleted                     = "customer.discount.deleted"            // Occurs whenever a customer's discount is removed.
	CustomerDiscountUpdated                     = "customer.discount.updated"            // Occurs whenever a customer is switched from one coupon to another.
	CustomerSourceCreated                       = "customer.source.created"              // Occurs whenever a new source is created for the customer.
	CustomerSourceDeleted                       = "customer.source.deleted"              // Occurs whenever a source is removed from a customer.
	CustomerSourceUpdated                       = "customer.source.updated"              // Occurs whenever a source's details are changed.
	CustomerSubscriptionCreated                 = "customer.subscription.created"        // Occurs whenever a customer with no subscription is signed up for a plan.
	CustomerSubscriptionDeleted                 = "customer.subscription.deleted"        // Occurs whenever a customer ends their subscription.
	CustomerSubscriptionTrialWillEnd            = "customer.subscription.trial_will_end" // Occurs three days before the trial period of a subscription is scheduled to end.
	CustomerSubscriptionUpdated                 = "customer.subscription.updated"        // Occurs whenever a subscription changes. Examples would include switching from one plan to another, or switching status from trial to active.
	InvoiceCreated                              = "invoice.created"                      // Occurs whenever a new invoice is created. If you are using webhooks, Stripe will wait one hour after they have all succeeded to attempt to pay the invoice; the only exception here is on the first invoice, which gets created and paid immediately when you subscribe a customer to a plan. If your webhooks do not all respond successfully, Stripe will continue retrying the webhooks every hour and will not attempt to pay the invoice. After 3 days, Stripe will attempt to pay the invoice regardless of whether or not your webhooks have succeeded. See  how to respond to a webhook.
	InvoicePaymentFailed                        = "invoice.payment_failed"               // Occurs whenever an invoice attempts to be paid, and the payment fails. This can occur either due to a declined payment, or because the customer has no active card. A particular case of note is that if a customer with no active card reaches the end of its free trial, an invoice.payment_failed notification will occur.
	InvoicePaymentSucceeded                     = "invoice.payment_succeeded"            // Occurs whenever an invoice attempts to be paid, and the payment succeeds.
	InvoiceUpdated                              = "invoice.updated"                      // Occurs whenever an invoice changes (for example, the amount could change).
	InvoiceitemCreated                          = "invoiceitem.created"                  // Occurs whenever an invoice item is created.
	InvoiceitemDeleted                          = "invoiceitem.deleted"                  // Occurs whenever an invoice item is deleted.
	InvoiceitemUpdated                          = "invoiceitem.updated"                  // Occurs whenever an invoice item is updated.
	OrderCreated                                = "order.created"                        // Occurs whenever an order is created.
	OrderPaymentFailed                          = "order.payment_failed"                 // Occurs whenever payment is attempted on an order, and the payment fails.
	OrderPaymentSucceeded                       = "order.payment_succeeded"              // Occurs whenever payment is attempted on an order, and the payment succeeds.
	OrderUpdated                                = "order.updated"                        // Occurs whenever an order is updated.
	OrderReturnCreated                          = "order_return.created"                 // Occurs whenever an order return created.
	PlanCreated                                 = "plan.created"                         // Occurs whenever a plan is created.
	PlanDeleted                                 = "plan.deleted"                         // Occurs whenever a plan is deleted.
	PlanUpdated                                 = "plan.updated"                         // Occurs whenever a plan is updated.
	ProductCreated                              = "product.created"                      // Occurs whenever a product is created.
	ProductDeleted                              = "product.deleted"                      // Occurs whenever a product is deleted.
	ProductUpdated                              = "product.updated"                      // Occurs whenever a product is updated.
	RecipientCreated                            = "recipient.created"                    // Occurs whenever a recipient is created.
	RecipientDeleted                            = "recipient.deleted"                    // Occurs whenever a recipient is deleted.
	RecipientUpdated                            = "recipient.updated"                    // Occurs whenever a recipient is updated.
	SkuCreated                                  = "sku.created"                          // Occurs whenever a SKU is created.
	SkuDeleted                                  = "sku.deleted"                          // Occurs whenever a SKU is deleted.
	SkuUpdated                                  = "sku.updated"                          // Occurs whenever a SKU is updated.
	TransferCreated                             = "transfer.created"                     // Occurs whenever a new transfer is created.
	TransferFailed                              = "transfer.failed"                      // Occurs whenever Stripe attempts to send a transfer and that transfer fails.
	TransferPaid                                = "transfer.paid"                        // Occurs whenever a sent transfer is expected to be available in the destination bank account. If the transfer failed, a transfer.failed webhook will additionally be sent at a later time. Note to Connect users: this event is only created for transfers from your connected Stripe accounts to their bank accounts, not for transfers to the connected accounts themselves.
	TransferReversed                            = "transfer.reversed"                    // Occurs whenever a transfer is reversed, including partial reversals.
	TransferUpdated                             = "transfer.updated"                     // Occurs whenever the description or metadata of a transfer is updated.
	Ping                                        = "ping"                                 // May be sent by Stripe at any time to see if a provided webhook URL is working.
)
