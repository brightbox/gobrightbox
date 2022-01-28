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
	assert.Equal(t, "img-3ikco", instance.ID, "image id incorrect")
}

func TestImage(t *testing.T) {
	instance := testInstance[Image](
		t,
		"Image",
		"images",
		"image",
		"img-3ikco",
	)
	assert.Equal(t, "img-3ikco", instance.ID, "image id incorrect")
	assert.Equal(t, "Ubuntu Lucid 10.04 server", instance.Name, "image name incorrect")
}

func TestCreateImage(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "POST",
			ExpectURL:    "/1.0/images",
			ExpectBody:   "{}",
			GiveBody:     readJSON("image"),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	newAC := ImageOptions{}
	ac, err := Create[Image](client, &newAC)
	assert.Assert(t, is.Nil(err), "Create[Image] returned an error")
	assert.Assert(t, ac != nil, "Create[Image] returned nil")
	assert.Equal(t, "img-3ikco", ac.ID)
}

func TestCreateImageWithSource(t *testing.T) {
	ts, client, err := SetupConnection(
		&APIMock{
			T:            t,
			ExpectMethod: "POST",
			ExpectURL:    "/1.0/images",
			ExpectBody:   `{"source":"ubuntu-lucid-daily-i64-server-20110509"}`,
			GiveBody:     readJSON("image"),
		},
	)
	defer ts.Close()
	assert.Assert(t, is.Nil(err), "Connect returned an error")

	pg := "ubuntu-lucid-daily-i64-server-20110509"
	newAC := ImageOptions{Source: &pg}
	ac, err := Create[Image](client, &newAC)
	assert.Assert(t, is.Nil(err), "CreateImage() returned an error")
	assert.Assert(t, ac != nil, "CreateImage() returned nil")
	assert.Equal(t, "img-3ikco", ac.ID)
	assert.Equal(t, pg, ac.Source)
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
