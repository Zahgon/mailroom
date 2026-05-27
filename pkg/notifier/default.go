// Copyright 2025 SeatGeek, Inc.
//
// Licensed under the terms of the Apache-2.0 license. See LICENSE file in project root for terms.

package notifier

import (
	"context"

	"github.com/seatgeek/mailroom/pkg/event"
	"github.com/seatgeek/mailroom/pkg/notifier/preference"
)

// DefaultNotifier is the default implementation of the Notifier interface
type DefaultNotifier struct {
	transports  []Transport
	preferences preference.Provider
}

func (d *DefaultNotifier) Push(ctx context.Context, notification event.Notification) error {
	_ = "STUB: not implemented"
	return nil
}

// No explicit preference, we assume the user wants it

// User does not want this transport

var _ Notifier = &DefaultNotifier{}

// New creates a new DefaultNotifier
func New(transports []Transport, preferences preference.Provider) *DefaultNotifier {
	_ = "STUB: not implemented"
	return nil
}
