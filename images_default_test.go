// Code generated by go generate; DO NOT EDIT.

package brightbox

import (
	"path"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestImages(t *testing.T) {
	instance := testAll(
		t,
		(*Client).Images,
		"Image",
		"images",
		"Images",
	)
	assert.Equal(t, instance.ID, "img-3ikco")
}

func TestImage(t *testing.T) {
	instance := testInstance(
		t,
		(*Client).Image,
		"Image",
		path.Join("images", "img-3ikco"),
		"image",
		"img-3ikco",
	)
	assert.Equal(t, instance.ID, "img-3ikco")
}

func TestCreateImage(t *testing.T) {
	newResource := ImageOptions{}
	instance := testModify(
		t,
		(*Client).CreateImage,
		newResource,
		"image",
		"POST",
		path.Join("images"),
		"{}",
	)
	assert.Equal(t, instance.ID, "img-3ikco")
}

func TestUpdateImage(t *testing.T) {
	updatedResource := ImageOptions{ID: "img-3ikco"}
	instance := testModify(
		t,
		(*Client).UpdateImage,
		updatedResource,
		"image",
		"PUT",
		path.Join("images", updatedResource.ID),
		"{}",
	)
	assert.Equal(t, instance.ID, updatedResource.ID)
}

func TestDestroyImage(t *testing.T) {
	deletedResource := testModify(
		t,
		(*Client).DestroyImage,
		"img-3ikco",
		"image",
		"DELETE",
		path.Join("images", "img-3ikco"),
		"",
	)
	assert.Equal(t, deletedResource.ID, "img-3ikco")
}

func TestLockImage(t *testing.T) {
	lockedResource := testModify(
		t,
		(*Client).LockImage,
		"img-3ikco",
		"image",
		"PUT",
		path.Join("images", "img-3ikco", "lock_resource"),
		"",
	)
	assert.Equal(t, lockedResource.ID, "img-3ikco")
}

func TestUnlockImage(t *testing.T) {
	unlockedResource := testModify(
		t,
		(*Client).UnlockImage,
		"img-3ikco",
		"image",
		"PUT",
		path.Join("images", "img-3ikco", "unlock_resource"),
		"",
	)
	assert.Equal(t, unlockedResource.ID, "img-3ikco")
}

func TestImageCreatedAtUnix(t *testing.T) {
	tm := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	target := Image{CreatedAt: &tm}
	assert.Equal(t, target.CreatedAtUnix(), tm.Unix())
}