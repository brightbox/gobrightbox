// Code generated by generate_enum; DO NOT EDIT.

package storagetype

import "fmt"

// Status is an enumerated type
type Status uint8

const (
	// Local is an enumeration for storagetype.Status
	Local Status = iota + 1
	// Network is an enumeration for Status
	Network
)

// Set of strings that are valid inputs for ParseStatus
var ValidStrings = []string{
	Local.String(),
	Network.String(),
}

// String makes Status satisfy the Stringer interface
func (i Status) String() string {
	tmp, err := i.MarshalText()
	if err == nil {
		return string(tmp)
	}
	return ""
}

// ParseStatus attempts to convert a string into a Status
func ParseStatus(name string) (Status, error) {
	switch name {
	case "local":
		return Local, nil
	case "network":
		return Network, nil
	}
	return Status(0), fmt.Errorf("%s is not a valid storagetype.Status", name)
}

// MarshalText implements the text marshaller method
func (i Status) MarshalText() ([]byte, error) {
	switch i {
	case Local:
		return []byte("local"), nil
	case Network:
		return []byte("network"), nil
	}
	return nil, fmt.Errorf("%v is not a valid storagetype.Status", i)
}

// UnmarshalText implements the text unmarshaller method
func (i *Status) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseStatus(name)
	if err != nil {
		return err
	}
	*i = tmp
	return nil
}