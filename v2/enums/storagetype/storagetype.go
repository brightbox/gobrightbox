// Code generated by generate_enum; DO NOT EDIT.

package storagetype

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Enum is an enumerated type
type Enum uint8

const (
	// Local is an enumeration for storagetype.Enum
	Local Enum = iota + 1
	// Network is an enumeration for storagetype.Enum
	Network
)

// Set of strings that are valid inputs for ParseEnum
var ValidStrings = []string{
	Local.String(),
	Network.String(),
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
	case "local":
		return Local, nil
	case "network":
		return Network, nil
	}
	var zero Enum
	return zero, fmt.Errorf("%s is not a valid storagetype.Enum", name)
}

// MarshalText implements the text marshaller method
func (i Enum) MarshalText() ([]byte, error) {
	switch i {
	case Local:
		return []byte("local"), nil
	case Network:
		return []byte("network"), nil
	}
	return nil, fmt.Errorf("%d is not a valid storagetype.Enum", i)
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