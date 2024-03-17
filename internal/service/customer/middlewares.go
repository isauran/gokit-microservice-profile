package customer

import (
	"context"
	"log/slog"
	"time"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(Service) Service

func LoggingMiddleware(l *slog.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next: next,
			log:  l,
		}
	}
}

type loggingMiddleware struct {
	next Service
	log  *slog.Logger
}

func (mw loggingMiddleware) PostCustomer(ctx context.Context, p Customer) (err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "PostCustomer", "id", p.ID, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.PostCustomer(ctx, p)
}

func (mw loggingMiddleware) GetCustomer(ctx context.Context, id string) (p Customer, err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "GetCustomer", "id", id, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.GetCustomer(ctx, id)
}

func (mw loggingMiddleware) PutCustomer(ctx context.Context, id string, p Customer) (err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "PutCustomer", "id", id, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.PutCustomer(ctx, id, p)
}

func (mw loggingMiddleware) PatchCustomer(ctx context.Context, id string, p Customer) (err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "PatchCustomer", "id", id, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.PatchCustomer(ctx, id, p)
}

func (mw loggingMiddleware) DeleteCustomer(ctx context.Context, id string) (err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "DeleteCustomer", "id", id, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.DeleteCustomer(ctx, id)
}

func (mw loggingMiddleware) GetAddresses(ctx context.Context, customerID string) (addresses []Address, err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "GetAddresses", "customerID", customerID, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.GetAddresses(ctx, customerID)
}

func (mw loggingMiddleware) GetAddress(ctx context.Context, customerID string, addressID string) (a Address, err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "GetAddress", "customerID", customerID, "addressID", addressID, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.GetAddress(ctx, customerID, addressID)
}

func (mw loggingMiddleware) PostAddress(ctx context.Context, customerID string, a Address) (err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "PostAddress", "customerID", customerID, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.PostAddress(ctx, customerID, a)
}

func (mw loggingMiddleware) DeleteAddress(ctx context.Context, customerID string, addressID string) (err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "DeleteAddress", "customerID", customerID, "addressID", addressID, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.DeleteAddress(ctx, customerID, addressID)
}
