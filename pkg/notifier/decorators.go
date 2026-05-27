// Copyright 2025 SeatGeek, Inc.
//
// Licensed under the terms of the Apache-2.0 license. See LICENSE file in project root for terms.

package notifier

import (
	"context"
	"log/slog"
	"time"

	"github.com/cenkalti/backoff/v5"
	"github.com/seatgeek/mailroom/pkg/event"
)

// WithTimeout decorates the given Transport with a timeout
func WithTimeout(transport Transport, timeout time.Duration) Transport {
	_ = "STUB: not implemented"
	return *new(Transport)
}

type withTimeout struct {
	Transport
	timeout time.Duration
}

func (w *withTimeout) Push(ctx context.Context, notification event.Notification) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *withTimeout) Validate(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

type (
	BackOff        = backoff.BackOff
	BackOffFactory = func() BackOff
)

// WithRetry decorates the given Transport with retry logic using the provided backoff
func WithRetry(transport Transport, maxTries uint, backoffFactory BackOffFactory) Transport {
	_ = "STUB: not implemented"
	return *new(Transport)
}

type withRetry struct {
	Transport
	maxTries uint
	backoff  func() BackOff
}

func (w *withRetry) Push(ctx context.Context, notification event.Notification) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *withRetry) Validate(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// WithLogging decorates the given Transport and logs every successful push (including the message body)
func WithLogging(transport Transport, logger *slog.Logger, logLevel slog.Level) Transport {
	_ = "STUB: not implemented"
	return *new(Transport)
}

type withLogging struct {
	Transport
	logger *slog.Logger
	level  slog.Level
}

func (w *withLogging) Push(ctx context.Context, n event.Notification) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *withLogging) Validate(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
