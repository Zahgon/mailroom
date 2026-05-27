// Copyright 2025 SeatGeek, Inc.
//
// Licensed under the terms of the Apache-2.0 license. See LICENSE file in project root for terms.

package user

import (
	"context"

	"github.com/seatgeek/mailroom/pkg/event"
)

// IdentifierEnrichmentProcessor is a processor that enriches notifications
// with all known identifiers for the recipient from the user.Store.
type IdentifierEnrichmentProcessor struct {
	userStore Store
}

// NewIdentifierEnrichmentProcessor creates a new IdentifierEnrichmentProcessor.
func NewIdentifierEnrichmentProcessor(us Store) *IdentifierEnrichmentProcessor {
	_ = "STUB: not implemented"
	return nil
}

// Process enriches each notification's recipient with additional identifiers.
func (p *IdentifierEnrichmentProcessor) Process(ctx context.Context, evt event.Event, notifications []event.Notification) ([]event.Notification, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Attempt to find the user based on the existing recipient identifiers.
