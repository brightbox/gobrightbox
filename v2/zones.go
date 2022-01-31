package brightbox

// Zone represents a Zone
// https://api.gb1.brightbox.com/1.0/#zone
type Zone struct {
	ID     string
	Handle string
}

// APIPath returns the relative URL path to the collection endpoint
func (c Zone) APIPath() string {
	return "zones"
}

// FetchID returns the ID field from the object
func (c Zone) FetchID() string {
	return c.ID
}

// HandleString returns the Handle field from a Zone object
func (c Zone) HandleString() string {
	return c.Handle
}
