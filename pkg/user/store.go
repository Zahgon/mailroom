// Copyright 2025 SeatGeek, Inc.
//
// Licensed under the terms of the Apache-2.0 license. See LICENSE file in project root for terms.

package user

import (
	"context"
	"errors"
	"sync"

	"github.com/seatgeek/mailroom/pkg/identifier"
	"github.com/seatgeek/mailroom/pkg/notifier/preference"
)

// ErrUserNotFound is returned when a user is not found in the Store.
// It should only be returned when the lookup process itself succeeded (e.g. no db connection errors), but
// we failed to locate a single known user in the store. Think of it like a 404 error.
var ErrUserNotFound = errors.New("user not found")

// Store is a database that stores user information, like Provider and identifiers.
// Implementations may be backed by a SQL database, an in-memory store, or something else.
//
// For all methods that search by identifier, the store MUST return the user that matches the identifier exactly.
// If no user matches the exact identifier, and the identifier is an "email" kind, the store SHOULD attempt to find
// a user where any email identifier matches the given email, regardless of namespace. This will allow for onboarding
// new integrations that utilize email identifiers without having to update all existing user information in the store.
type Store interface {
	// Get returns a user by its key, or an error if the user is not found
	Get(ctx context.Context, key string) (*User, error)
	// GetByIdentifier returns a user by a given identifier, or an error if the user is not found
	GetByIdentifier(ctx context.Context, identifier identifier.Identifier) (*User, error)

	// Find searches for a user matching any of the given identifiers
	// (The user is not required to match all of them, just one is enough)
	Find(ctx context.Context, possibleIdentifiers identifier.Set) (*User, error)

	// SetPreferences replaces the preferences for a user by key
	SetPreferences(ctx context.Context, key string, prefs preference.Map) error
}

// InMemoryStore is a simple in-memory implementation of the Store interface
// This is especially useful for testing, but can also be used for simple applications which don't need durable preference storage.
type InMemoryStore struct {
	users []*User
	mu    sync.RWMutex
}

var _ Store = &InMemoryStore{}

// NewInMemoryStore creates a new in-memory store with the given users
func NewInMemoryStore(users ...*User) *InMemoryStore { _ = "STUB: not implemented"; return nil }

// Add adds a user to the in-memory store
func (s *InMemoryStore) Add(_ context.Context, u *User) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *InMemoryStore) Get(_ context.Context, key string) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *InMemoryStore) GetByIdentifier(ctx context.Context, identifier identifier.Identifier) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Look for an exact match

// Or if the identifier is an email, look for any matching email

func (s *InMemoryStore) Find(ctx context.Context, possibleIdentifiers identifier.Set) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *InMemoryStore) SetPreferences(ctx context.Context, key string, prefs preference.Map) error {
	_ = "STUB: not implemented"
	return nil
}
