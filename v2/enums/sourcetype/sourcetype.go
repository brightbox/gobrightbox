// Code generated by generate_enum; DO NOT EDIT.

package sourcetype

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Enum is an enumerated type
type Enum uint8

const (
	// Upload is an enumeration for sourcetype.Enum
	Upload Enum = iota + 1
	// Snapshot is an enumeration for sourcetype.Enum
	Snapshot
)

// Set of strings that are valid inputs for ParseEnum
var ValidStrings = []string{
	Upload.String(),
	Snapshot.String(),
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
	case "upload":
		return Upload, nil
	case "snapshot":
		return Snapshot, nil
	}
	var zero Enum
	return zero, fmt.Errorf("%s is not a valid sourcetype.Enum", name)
}

// MarshalText implements the text marshaller method
func (i Enum) MarshalText() ([]byte, error) {
	switch i {
	case Upload:
		return []byte("upload"), nil
	case Snapshot:
		return []byte("snapshot"), nil
	}
	return nil, fmt.Errorf("%d is not a valid sourcetype.Enum", i)
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