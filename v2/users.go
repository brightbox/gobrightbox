package brightbox

import "time"

// User represents a Brightbox User
// https://api.gb1.brightbox.com/1.0/#user
type User struct {
	ResourceRef
	ID            string
	Name          string
	EmailAddress  string     `json:"email_address"`
	EmailVerified bool       `json:"email_verified"`
	SSHKey        string     `json:"ssh_key"`
	MessagingPref bool       `json:"messaging_pref"`
	CreatedAt     *time.Time `json:"created_at"`
	TwoFactorAuth struct {
		Enabled bool
	} `json:"2fa"`
	Accounts       []Account
	DefaultAccount *Account `json:"default_account"`
}

// UserOptions is used to update objects
type UserOptions struct {
	ID                   string  `json:"-"`
	Name                 *string `json:"name,omitempty"`
	EmailAddress         *string `json:"email_address,omitempty"`
	SSHKey               *string `json:"ssh_key,omitempty"`
	Password             *string `json:"password,omitempty"`
	PasswordConfirmation *string `json:"password_confirmation,omitempty"`
}

// APIPath returns the relative URL path to the collection endpoint
func (c User) APIPath() string {
	return "users"
}

// FetchID returns the ID field from the object
func (c User) FetchID() string {
	return c.ID
}

// PutPath returns the relative URL path to PUT an object
func (c User) PutPath(from *UserOptions) string {
	return c.APIPath() + "/" + from.OptionID()
}

// OptionID returns the ID field from and options object
// ID will be blank for create, and set for update
func (c UserOptions) OptionID() string {
	return c.ID
}
