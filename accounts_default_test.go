// Code generated by go generate; DO NOT EDIT.

package brightbox

import (
	"path"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestAccounts(t *testing.T) {
	instance := testAll(
		t,
		(*Client).Accounts,
		"Account",
		"accounts",
		"Accounts",
	)
	assert.Equal(t, instance.ID, "acc-43ks4")
}

func TestAccount(t *testing.T) {
	instance := testInstance(
		t,
		(*Client).Account,
		"Account",
		path.Join("accounts", "acc-43ks4"),
		"account",
		"acc-43ks4",
	)
	assert.Equal(t, instance.ID, "acc-43ks4")
}

func TestUpdateAccount(t *testing.T) {
	updatedResource := AccountOptions{ID: "acc-43ks4"}
	instance := testModify(
		t,
		(*Client).UpdateAccount,
		updatedResource,
		"account",
		"PUT",
		path.Join("accounts", updatedResource.ID),
		"{}",
	)
	assert.Equal(t, instance.ID, updatedResource.ID)
}

func TestAccountCreatedAtUnix(t *testing.T) {
	tm := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	target := Account{CreatedAt: &tm}
	assert.Equal(t, target.CreatedAtUnix(), tm.Unix())
}

func TestResetAccountPassword(t *testing.T) {
	instance := testModify(
		t,
		(*Client).ResetAccountPassword,
		"acc-43ks4",
		"account",
		"POST",
		path.Join("accounts", "acc-43ks4", "reset_ftp_password"),
		"",
	)
	assert.Equal(t, instance.ID, "acc-43ks4")
}