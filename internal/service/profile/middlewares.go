package profile

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

func (mw loggingMiddleware) PostProfile(ctx context.Context, p Profile) (err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "PostProfile", "id", p.ID, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.PostProfile(ctx, p)
}

func (mw loggingMiddleware) GetProfile(ctx context.Context, id string) (p Profile, err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "GetProfile", "id", id, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.GetProfile(ctx, id)
}

func (mw loggingMiddleware) PutProfile(ctx context.Context, id string, p Profile) (err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "PutProfile", "id", id, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.PutProfile(ctx, id, p)
}

func (mw loggingMiddleware) PatchProfile(ctx context.Context, id string, p Profile) (err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "PatchProfile", "id", id, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.PatchProfile(ctx, id, p)
}

func (mw loggingMiddleware) DeleteProfile(ctx context.Context, id string) (err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "DeleteProfile", "id", id, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.DeleteProfile(ctx, id)
}

func (mw loggingMiddleware) GetAddresses(ctx context.Context, profileID string) (addresses []Address, err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "GetAddresses", "profileID", profileID, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.GetAddresses(ctx, profileID)
}

func (mw loggingMiddleware) GetAddress(ctx context.Context, profileID string, addressID string) (a Address, err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "GetAddress", "profileID", profileID, "addressID", addressID, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.GetAddress(ctx, profileID, addressID)
}

func (mw loggingMiddleware) PostAddress(ctx context.Context, profileID string, a Address) (err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "PostAddress", "profileID", profileID, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.PostAddress(ctx, profileID, a)
}

func (mw loggingMiddleware) DeleteAddress(ctx context.Context, profileID string, addressID string) (err error) {
	defer func(begin time.Time) {
		mw.log.Debug("method", "DeleteAddress", "profileID", profileID, "addressID", addressID, "took", time.Since(begin).String(), "err", err)
	}(time.Now())
	return mw.next.DeleteAddress(ctx, profileID, addressID)
}
