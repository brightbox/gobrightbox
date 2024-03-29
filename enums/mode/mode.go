// Code generated by generate_enum; DO NOT EDIT.

// Package mode is an enumeration of the states Nat, Route
package mode

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Enum is an enumerated type
type Enum uint8

const (
	// Nat is an enumeration for mode.Enum
	Nat Enum = iota + 1
	// Route is an enumeration for mode.Enum
	Route
)

// ValidStrings is the set of strings that are valid inputs to ParseEnum
var ValidStrings = []string{
	Nat.String(),
	Route.String(),
}

// String makes Enum satisfy the Stringer interface
func (i Enum) String() string {
	tmp, err := i.MarshalText()
	if err == nil {
		return string(tmp)
	}
	return ""
}

// ParseEnum attempts to convert a string into a Enum
func ParseEnum(name string) (Enum, error) {
	switch name {
	case "nat":
		return Nat, nil
	case "route":
		return Route, nil
	}
	var zero Enum
	return zero, fmt.Errorf("%s is not a valid mode.Enum", name)
}

// MarshalText implements the text marshaller method
func (i Enum) MarshalText() ([]byte, error) {
	switch i {
	case Nat:
		return []byte("nat"), nil
	case Route:
		return []byte("route"), nil
	}
	return nil, fmt.Errorf("%d is not a valid mode.Enum", i)
}

// UnmarshalText implements the text unmarshaller method
func (i *Enum) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseEnum(name)
	if err != nil {
		return &json.UnmarshalTypeError{
			Value: name,
			Type:  reflect.TypeOf(*i),
		}
	}
	*i = tmp
	return nil
}
