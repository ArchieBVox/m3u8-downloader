package tool

import (
	"io"
	"testing"
)

func TestGet(t *testing.T) {
	body, err := Get("https://raw.githubusercontent.com/ArchieBVox/m3u8-downloader/develop/README.md")
	if err != nil {
		t.Error(err)
	}
	defer body.Close()

	_, err = io.ReadAll(body)
	if err != nil {
		t.Error(err)
	}
}
