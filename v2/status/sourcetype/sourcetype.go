// Code generated by generate_enum; DO NOT EDIT.

package sourcetype

import "fmt"

// Status is an enumerated type
type Status uint8

const (
	// Upload is an enumeration for sourcetype.Status
	Upload Status = iota + 1
	// Snapshot is an enumeration for Status
	Snapshot
)

// Set of strings that are valid inputs for ParseStatus
var ValidStrings = []string{
	Upload.String(),
	Snapshot.String(),
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
	case "upload":
		return Upload, nil
	case "snapshot":
		return Snapshot, nil
	}
	return Status(0), fmt.Errorf("%s is not a valid sourcetype.Status", name)
}

// MarshalText implements the text marshaller method
func (i Status) MarshalText() ([]byte, error) {
	switch i {
	case Upload:
		return []byte("upload"), nil
	case Snapshot:
		return []byte("snapshot"), nil
	}
	return nil, fmt.Errorf("%v is not a valid sourcetype.Status", i)
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