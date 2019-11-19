package test

import (
	"encoding/base64"
	"io"
	"strings"
)

func (h TemplateA) Template() string {
	result := &strings.Builder{}
	if _, err := io.Copy(result, base64.NewDecoder(base64.StdEncoding, strings.NewReader("cGFja2FnZSBhdXRvCgoKZnVuYyB7ey5OYW1lfX0oKSBpbnQgIHsKICAgIHJldHVybiB7ey5WYWx1ZX19Cn0K"))); err != nil {
		panic(err)
	}
	return result.String()
}
