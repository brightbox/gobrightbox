package brightbox

import (
	"testing"

	"gotest.tools/assert"
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
	name := "dev client"
	uac := ImageOptions{ID: "img-3ikco", Name: &name}
	_ = testUpdate[Image](
		t,
		"Image",
		"images",
		"image",
		"img-3ikco",
		&uac,
		`{"name":"dev client"}`,
	)
}

func TestDestroyImage(t *testing.T) {
	testDestroy[Image](
		t,
		"Image",
		"images",
		"img-3ikco",
	)
}

func TestLockImage(t *testing.T) {
	testLock[Image](
		t,
		"Image",
		"images",
		&Image{ID: "img-3ikco"},
		"img-3ikco",
		"lock_resource",
		LockResource,
	)
}

func TestUnlockImage(t *testing.T) {
	testLock[Image](
		t,
		"Image",
		"images",
		&Image{ID: "img-3ikco"},
		"img-3ikco",
		"unlock_resource",
		UnlockResource,
	)
}
