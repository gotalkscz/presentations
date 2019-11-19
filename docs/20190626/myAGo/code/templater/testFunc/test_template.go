package testFunc

import (
	"encoding/base64"
	"io"
	"strings"
)

func (h TemplateA) Template() string {
	result := &strings.Builder{}
	if _, err := io.Copy(result, base64.NewDecoder(base64.StdEncoding, strings.NewReader("cGFja2FnZSBhdXRvCgoKZnVuYyB7ey5OYW1lfX0oKSBzdHJpbmcgIHsKICAgIHJldHVybiB7e0dvVmFsdWUgLkZ1bGx9fQp9Cg=="))); err != nil {
		panic(err)
	}
	return result.String()
}
