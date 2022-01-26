package brightbox

type optionID interface{
	FetchID() string
}

type crudable[I optionID] interface {
	queriable
	Extract() *I
}

// Create creates a new resource from the supplied option map
//
// It takes an instance of Options. Not all attributes can be
// specified at create time (such as ID, which is allocated for you).
func Create[T crudable[I], I optionID](q *Client, newOptions *I) (*T, error) {
	var resource T
	_, err := q.MakeAPIRequest(
		"POST",
		resource.APIPath(),
		newOptions,
		&resource,
	)
	return &resource, err
}

// Update updates an existing resources's attributes. Not all
// attributes can be changed (such as ID).
//
// Specify the resource you want to update using the ID field
// field.
func Update[T crudable[I], I optionID](q *Client, updateOptions *I) (*T, error) {
	var resource T
	_, err := q.MakeAPIRequest(
		"PUT",
		resource.APIPath() + "/" + (*updateOptions).FetchID(),
		updateOptions,
		&resource,
	)
	return &resource, err
}

// Destroy destroys an existing resource.
func Destroy[T queriable](q *Client, identifier string) error {
	var zero T
	_, err := q.MakeAPIRequest(
		"DELETE",
		zero.APIPath() + "/" + identifier,
		nil,
		nil,
	)
	return err
}
