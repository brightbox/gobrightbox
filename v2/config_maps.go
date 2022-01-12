package brightbox

// ConfigMap represents a config map
// https://api.gb1.brightbox.com/1.0/#config_maps
type ConfigMap struct {
	ID   string                 `json:"id"`
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}

// ConfigMapOptions is used in combination with CreateConfigMap and
// UpdateConfigMap to create and update config maps
type ConfigMapOptions struct {
	ID   string                  `json:"-"`
	Name *string                 `json:"name,omitempty"`
	Data *map[string]interface{} `json:"data,omitempty"`
}

func (_c ConfigMap) APIPath() string {
	return "config_maps"
}
