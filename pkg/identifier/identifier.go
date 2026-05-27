// Copyright 2025 SeatGeek, Inc.
//
// Licensed under the terms of the Apache-2.0 license. See LICENSE file in project root for terms.

// Package identifier provides a way to identify users across different systems.
package identifier

import (
	"sync"
)

// NamespaceAndKind is a combination of a namespace and a kind.
// For example, in "slack.com/email", "slack.com" is the namespace and "email" is the kind.
// The namespace part is considered optional, and if it is not present, it is represented as an empty string.
// This is useful when a user is known by different emails, usernames, or IDs across different systems.
type NamespaceAndKind string

var (
	// GenericEmail is any email address not associated with a specific namespace or system.
	GenericEmail = NamespaceAndKind(KindEmail)

	// GenericUsername is any username not associated with a specific namespace or system.
	GenericUsername = NamespaceAndKind(KindUsername)

	// GenericID is any ID not associated with a specific namespace or system.
	GenericID = NamespaceAndKind(KindID)
)

// Split returns the namespace and kind parts of the NamespaceAndKind.
func (n NamespaceAndKind) Split() (string, string) { _ = "STUB: not implemented"; return "", "" }

func (n NamespaceAndKind) Namespace() string { _ = "STUB: not implemented"; return "" }

func (n NamespaceAndKind) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

// NewNamespaceAndKind creates a new NamespaceAndKind from a namespace and a Kind.
func NewNamespaceAndKind[T ~string](namespace string, kind T) NamespaceAndKind {
	_ = "STUB: not implemented"
	return *new(NamespaceAndKind)
}

// Kind represents the type of identifier, such as an "email" or "username".
// This is used in conjunction with a namespace to uniquely identify a user in some system.
type Kind string

const (
	KindEmail    Kind = "email"
	KindUsername Kind = "username"
	KindID       Kind = "id"
)

// An Identifier is a unique reference to some user or group.
type Identifier struct {
	NamespaceAndKind
	Value string
}

// valueType is used by the generic New function to allow any string or integer type to be passed as the value argument.
type valueType interface {
	~string | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// New creates a new Identifier for a given namespaceAndKind and a value.
func New[T1 ~string, T2 valueType](namespaceAndKind T1, value T2) Identifier {
	_ = "STUB: not implemented"
	return *new(Identifier)
}

// Set holds a thread-safe map of NamespaceAndKind to a value.
// Each entry is basically an Identifier.
type Set interface {
	Get(NamespaceAndKind) (string, bool)
	MustGet(NamespaceAndKind) string
	Add(Identifier)
	Merge(Set)
	Intersect(Set) Set
	ToList() []Identifier
	String() string
	ToMap() map[NamespaceAndKind]string
	Len() int
	Copy() Set
}

type set struct {
	ids   map[NamespaceAndKind]string
	mutex sync.RWMutex
}

func (c *set) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *set) Get(namespaceAndKind NamespaceAndKind) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (c *set) MustGet(namespaceAndKind NamespaceAndKind) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *set) Add(id Identifier) { _ = "STUB: not implemented"; return }

// Merge adds all the identifiers from another Set to this Set.
func (c *set) Merge(otherIdentifiers Set) { _ = "STUB: not implemented"; return }

// Intersect returns a new Set that contains only the identifiers that are present in both this Set and another Set.
func (c *set) Intersect(other Set) Set { _ = "STUB: not implemented"; return *new(Set) }

// ToList returns the Set as a slice of Identifier objects.
func (c *set) ToList() []Identifier { _ = "STUB: not implemented"; return nil }

func (c *set) String() string { _ = "STUB: not implemented"; return "" }

func (c *set) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Copy creates a deep copy of the Set.
func (c *set) Copy() Set { _ = "STUB: not implemented"; return *new(Set) }

// NewSet creates a new Set from a slice of Identifier objects
func NewSet(ids ...Identifier) Set { _ = "STUB: not implemented"; return *new(Set) }

// NewSetFromMap creates a new Set from a map of NamespaceAndKind to value.
func NewSetFromMap(ids map[NamespaceAndKind]string) Set {
	_ = "STUB: not implemented"
	return *new(Set)
}

// ToMap returns the Set as a map of NamespaceAndKind to value from a Set.
func (c *set) ToMap() map[NamespaceAndKind]string { _ = "STUB: not implemented"; return nil }
