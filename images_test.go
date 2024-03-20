package brightbox

import (
	"testing"

	"github.com/brightbox/gobrightbox/v2/enums/arch"
	"gotest.tools/v3/assert"
)

func TestCreateImageWithSource(t *testing.T) {
	pg := "https://api.gb1.brightbox.com/1.0/images/img-3ikco"
	newResource := ImageOptions{URL: pg}
	instance := testModify(
		t,
		(*Client).CreateImage,
		newResource,
		"image",
		"POST",
		"images",
		`{"http_url":"https://api.gb1.brightbox.com/1.0/images/img-3ikco"}`,
	)
	assert.Equal(t, instance.URL, pg)
	assert.Equal(t, instance.Arch, arch.X86_64)
}
