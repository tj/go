// Package hooks provides Stripe webhook handling with secret support.
package hooks

import (
	"io/ioutil"
	"net/http"

	"github.com/apex/log"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/webhook"
	"github.com/tj/go/http/response"
)

// Func is an event handler function.
type Func func(*stripe.Event) error

// New event handler with secret.
func New(secret string, h Func) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.WithError(err).Error("reading body")
			response.InternalServerError(w)
			return
		}

		signature := r.Header.Get("Stripe-Signature")
		e, err := webhook.ConstructEvent(b, signature, secret)
		if err != nil {
			log.WithError(err).Error("constructing event")
			response.InternalServerError(w)
			return
		}

		ctx := log.WithFields(log.Fields{
			"event_id":   e.ID,
			"event_type": e.Type,
		})

		ctx.Info("handling stripe event")
		if err := h(&e); err != nil {
			ctx.WithError(err).Error("handling stripe event")
			response.InternalServerError(w)
			return
		}

		ctx.Info("handled stripe event")
		response.OK(w)
	})
}
