package brightbox

// ConfigMap represents a config map
// https://api.gb1.brightbox.com/1.0/#config_maps
type ConfigMap struct {
	ID   string                 `json:"id"`
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}

// ConfigMapOptions is used to create and update config maps
type ConfigMapOptions struct {
	ID   string                  `json:"-"`
	Name *string                 `json:"name,omitempty"`
	Data *map[string]interface{} `json:"data,omitempty"`
}

// APIPath returns the relative URL path to the collection endpoint
func (c ConfigMap) APIPath() string {
	return "config_maps"
}

// PostPath returns the relative URL path to POST an object
func (c ConfigMap) PostPath(from *ConfigMapOptions) string {
	return c.APIPath()
}

// PutPath returns the relative URL path to PUT an object
func (c ConfigMap) PutPath(from *ConfigMapOptions) string {
	return c.APIPath() + "/" + from.OptionID()
}

// DestroyPath returns the relative URL path to DESTROY an object
func (c ConfigMap) DestroyPath(from string) string {
	return c.APIPath() + "/" + from
}

// FetchID returns the ID field from the object
func (c ConfigMap) FetchID() string {
	return c.ID
}

// OptionID returns the ID field from and options object
// ID will be blank for create, and set for update
func (c ConfigMapOptions) OptionID() string {
	return c.ID
}
