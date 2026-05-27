// Copyright 2025 SeatGeek, Inc.
//
// Licensed under the terms of the Apache-2.0 license. See LICENSE file in project root for terms.

// Package postgres provides a postgresql-backed implementation of the user.Store interface
package postgres

import (
	"context"
	"time"

	"github.com/seatgeek/mailroom/pkg/identifier"
	"github.com/seatgeek/mailroom/pkg/notifier/preference"
	"github.com/seatgeek/mailroom/pkg/user"
	"gorm.io/gorm"
)

// UserModel is the gorm model for a user
type UserModel struct {
	Key         string         `gorm:"primarykey"`
	Preferences preference.Map `gorm:"serializer:json"`

	// Identifiers is a map of all identifiers for the user
	Identifiers map[identifier.NamespaceAndKind]string `gorm:"serializer:json"`
	// Emails contains the subset of Identifiers that have Kind=="email" (for easier fallback lookup)
	Emails []string `gorm:"serializer:json"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *UserModel) TableName() string {
	_ = "STUB: not implemented"

	// ToUser converts a UserModel to a user.User
	return ""
}

func (u *UserModel) ToUser() *user.User { _ = "STUB: not implemented"; return nil }

type Store struct {
	db *gorm.DB
}

// NewPostgresStore creates a new postgres store
func NewPostgresStore(db *gorm.DB) *Store { _ = "STUB: not implemented"; return nil }

// Add upserts a user to the postgres store
func (s *Store) Add(ctx context.Context, u *user.User) error { _ = "STUB: not implemented"; return nil }

// Find implements user.Store.
func (s *Store) Find(ctx context.Context, possibleIdentifiers identifier.Set) (*user.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No users found; fall back to email identifiers if possible

// Get implements user.Store.
func (s *Store) Get(ctx context.Context, key string) (*user.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetByIdentifier implements user.Store.
func (s *Store) GetByIdentifier(ctx context.Context, id identifier.Identifier) (*user.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fall back to any email identifier

// SetPreferences implements user.Store.
func (s *Store) SetPreferences(ctx context.Context, key string, prefs preference.Map) error {
	_ = "STUB: not implemented"
	return nil
}

var _ user.Store = &Store{}
