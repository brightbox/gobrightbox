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

// APIPath returns the relative URL path to the config map collection
func (c ConfigMap) APIPath() string {
	return "config_maps"
}

// Extract copies a ConfigMap object to a ConfigMapOptions object
func (c ConfigMap) Extract() *ConfigMapOptions {
	return &ConfigMapOptions{
		ID:   c.ID,
		Name: &c.Name,
		Data: &c.Data,
	}
}

// FetchID returns the ID field from a ConfigMapOptions object
// ID will be blank for create, and set for update
func (c ConfigMapOptions) FetchID() string {
	return c.ID
}
