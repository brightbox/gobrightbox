// Code generated by generate_enum; DO NOT EDIT.

package healthchecktype

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Enum is an enumerated type
type Enum uint8

const (
	// Tcp is an enumeration for healthchecktype.Enum
	Tcp Enum = iota + 1
	// Http is an enumeration for healthchecktype.Enum
	Http
)

// Set of strings that are valid inputs for ParseEnum
var ValidStrings = []string{
	Tcp.String(),
	Http.String(),
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
	case "tcp":
		return Tcp, nil
	case "http":
		return Http, nil
	}
	var zero Enum
	return zero, fmt.Errorf("%s is not a valid healthchecktype.Enum", name)
}

// MarshalText implements the text marshaller method
func (i Enum) MarshalText() ([]byte, error) {
	switch i {
	case Tcp:
		return []byte("tcp"), nil
	case Http:
		return []byte("http"), nil
	}
	return nil, fmt.Errorf("%d is not a valid healthchecktype.Enum", i)
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