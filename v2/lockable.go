package brightbox

type lockable interface {
	LockID() string
}

// LockResource locks a resource against destroy requests.
func LockResource(q *Client, resource lockable) error {
	_, err := q.MakeAPIRequest(
		"PUT",
		resource.LockID()+"/lock_resource",
		nil,
		nil,
	)
	return err
}

// UnlockResource unlocks a resource, renabling destroy requests.
func UnlockResource(q *Client, resource lockable) error {
	_, err := q.MakeAPIRequest(
		"PUT",
		resource.LockID()+"/unlock_resource",
		nil,
		nil,
	)
	return err
}
