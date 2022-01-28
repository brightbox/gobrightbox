package brightbox

import (
	"time"
)

// Account represents a Brightbox Cloud Account
// https://api.gb1.brightbox.com/1.0/#account
type Account struct {
	ID                    string
	Name                  string
	Status                string
	Address1              string `json:"address_1"`
	Address2              string `json:"address_2"`
	City                  string
	County                string
	Postcode              string
	CountryCode           string     `json:"country_code"`
	CountryName           string     `json:"country_name"`
	VatRegistrationNumber string     `json:"vat_registration_number"`
	TelephoneNumber       string     `json:"telephone_number"`
	TelephoneVerified     bool       `json:"telephone_verified"`
	VerifiedTelephone     string     `json:"verified_telephone"`
	VerifiedAt            *time.Time `json:"verified_at"`
	VerifiedIP            string     `json:"verified_ip"`
	ValidCreditCard       bool       `json:"valid_credit_card"`
	CreatedAt             *time.Time `json:"created_at"`
	RAMLimit              int        `json:"ram_limit"`
	RAMUsed               int        `json:"ram_used"`
	DbsRAMLimit           int        `json:"dbs_ram_limit"`
	DbsRAMUsed            int        `json:"dbs_ram_used"`
	CloudIPsLimit         int        `json:"cloud_ips_limit"`
	CloudIPsUsed          int        `json:"cloud_ips_used"`
	LoadBalancersLimit    int        `json:"load_balancers_limit"`
	LoadBalancersUsed     int        `json:"load_balancers_used"`
	LibraryFtpHost        string     `json:"library_ftp_host"`
	LibraryFtpUser        string     `json:"library_ftp_user"`
	LibraryFtpPassword    string     `json:"library_ftp_password"`
	Owner                 User
	Users                 []User
}

// APIPath returns the relative URL path to the accounts collection
func (c Account) APIPath() string {
	return "accounts"
}
