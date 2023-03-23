package echopongo2

import (
	"bytes"
	"embed"
	"io"
	"path"

	"github.com/flosch/pongo2/v4"
)

func NewLoader(prefix string, fs *embed.FS) pongo2.TemplateLoader {
	return &Loader{prefix: prefix, fs: fs}
}

type Loader struct {
	prefix string
	fs     *embed.FS
}

func (l *Loader) Abs(base, name string) string {
	return name
}

func (l *Loader) Get(name string) (io.Reader, error) {
	name = path.Join(l.prefix, name)
	b, err := l.fs.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
