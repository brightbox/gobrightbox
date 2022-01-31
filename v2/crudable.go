package brightbox

type optionID interface{
	OptionID() string
}

type createable[I optionID] interface {
	FetchID() string
	PostPath(from *I) string
}

type updateable[I optionID] interface {
	FetchID() string
	PutPath(from *I) string
}

type destroyable interface {
	DestroyPath(from string) string
}

// Create creates a new resource from the supplied option map
//
// It takes an instance of Options. Not all attributes can be
// specified at create time (such as ID, which is allocated for you).
func Create[T createable[I], I optionID](q *Client, newOptions *I) (*T, error) {
	var resource T
	_, err := q.MakeAPIRequest(
		"POST",
		resource.PostPath(newOptions),
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
func Update[T updateable[I], I optionID](q *Client, updateOptions *I) (*T, error) {
	var resource T
	_, err := q.MakeAPIRequest(
		"PUT",
		resource.PutPath(updateOptions),
		updateOptions,
		&resource,
	)
	return &resource, err
}

// Destroy destroys an existing resource.
func Destroy[T destroyable](q *Client, identifier string) error {
	var zero T
	_, err := q.MakeAPIRequest(
		"DELETE",
		zero.DestroyPath(identifier),
		nil,
		nil,
	)
	return err
}
