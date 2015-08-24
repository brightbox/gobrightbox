package brightbox

type Account struct {
	Resource
	Name                  string
	Status                string
	Address1              string `json:"address_1"`
	Address2              string `json:"address_2"`
	City                  string
	County                string
	Postcode              string
	CountryCode           string
	CountryName           string
	VatRegistrationNumber string `json:"vat_registration_number"`
	TelephoneNumber       string `json:"telephone_number"`
	TelephoneVerified     bool   `json:"telephone_verified"`
	VerifiedTelephone     string `json:"verified_telephone"`
	RamUsed               int    `json:"ram_used"`
}

func (c *Client) Accounts() (*[]Account, error) {
	accounts := new([]Account)
	_, err := c.MakeApiRequest("GET", "/1.0/accounts", nil, accounts)
	if err != nil {
		return nil, err
	}
	return accounts, err
}
