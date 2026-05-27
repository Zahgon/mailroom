// Copyright 2025 SeatGeek, Inc.
//
// Licensed under the terms of the Apache-2.0 license. See LICENSE file in project root for terms.

package user

import (
	"context"

	"github.com/seatgeek/mailroom/pkg/event"
	"github.com/seatgeek/mailroom/pkg/notifier/preference"
)

type PreferenceProvider struct {
	userStore Store
}

var _ preference.Provider = (*PreferenceProvider)(nil)

func NewPreferenceProvider(userStore Store) *PreferenceProvider {
	_ = "STUB: not implemented"
	return nil
}

func (p PreferenceProvider) Wants(ctx context.Context, notification event.Notification, transport event.TransportKey) *bool {
	_ = "STUB: not implemented"
	return nil
}

func (p PreferenceProvider) getRecipientUserForNotification(ctx context.Context, notification event.Notification) *User {
	_ = "STUB: not implemented"
	return nil
}
