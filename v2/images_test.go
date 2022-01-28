package brightbox

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestImages(t *testing.T) {
	instance := testAll[Image](
		t,
		"Image",
		"images",
		"image",
	)
	assert.Equal(t, instance.ID, "img-3ikco")
}

func TestImage(t *testing.T) {
	instance := testInstance[Image](
		t,
		"Image",
		"images",
		"image",
		"img-3ikco",
	)
	assert.Equal(t, instance.ID, "img-3ikco")
	assert.Equal(t, instance.Name, "Ubuntu Lucid 10.04 server")
}

func TestCreateImage(t *testing.T) {
	newAC := ImageOptions{}
	_ = testCreate[Image](
		t,
		"Image",
		"images",
		"image",
		"img-3ikco",
		&newAC,
		"{}",
	)
}

func TestCreateImageWithSource(t *testing.T) {
	pg := "ubuntu-lucid-daily-i64-server-20110509"
	newAC := ImageOptions{Source: &pg}
	instance := testCreate[Image](
		t,
		"Image",
		"images",
		"image",
		"img-3ikco",
		&newAC,
		`{"source":"ubuntu-lucid-daily-i64-server-20110509"}`,
	)
	assert.Equal(t, instance.Source, pg)
}

func TestUpdateImage(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "PUT",
			ExpectURL:    "/1.0/images/img-3ikco",
			ExpectBody:   `{"name":"dev client"}`,
			GiveBody:     readJSON("image"),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	name := "dev client"
	uac := ImageOptions{ID: "img-3ikco", Name: &name}
	ac, err := Update[Image](client, &uac)
	assert.Assert(t, is.Nil(err), "UpdateImage() returned an error")
	assert.Assert(t, ac != nil, "UpdateImage() returned nil")
	assert.Equal(t, "img-3ikco", ac.ID)
}

func TestDestroyImage(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "DELETE",
			ExpectURL:    "/1.0/images/img-3ikco",
			ExpectBody:   "",
			GiveBody:     "",
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	err = Destroy[Image](client, "img-3ikco")
	assert.Assert(t, is.Nil(err), "DestroyImage() returned an error")
}

func TestLockImage(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "PUT",
			ExpectURL:    "/1.0/images/img-3ikco/lock_resource",
			ExpectBody:   ``,
			GiveBody:     ``,
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	err = LockResource(client, &Image{ID: "img-3ikco"})
	assert.Assert(t, is.Nil(err), "LockImage() returned an error")
}

func TestUnlockImage(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "PUT",
			ExpectURL:    "/1.0/images/img-3ikco/unlock_resource",
			ExpectBody:   ``,
			GiveBody:     ``,
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	err = UnLockResource(client, &Image{ID: "img-3ikco"})
	assert.Assert(t, is.Nil(err), "LockImage() returned an error")
}
