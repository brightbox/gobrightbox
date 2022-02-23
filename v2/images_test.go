package brightbox

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestCreateImageWithSource(t *testing.T) {
	pg := "ubuntu-lucid-daily-i64-server-20110509"
	newResource := ImageOptions{Source: &pg}
	instance := testModify[Image, *ImageOptions](
		t,
		(*Client).CreateImage,
		&newResource,
		"image",
		"POST",
		"images",
		`{"source":"ubuntu-lucid-daily-i64-server-20110509"}`,
	)
	assert.Equal(t, instance.Source, pg)
}
