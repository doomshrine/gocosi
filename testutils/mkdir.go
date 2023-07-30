package testutils

import (
	"fmt"
	"net/url"
	"os"
	"path"

	"github.com/doomshrine/must"
)

func MustMkdirTemp() string {
	return must.Do(os.MkdirTemp("", "*"))
}

func MustMkUnixTemp(socketName string) *url.URL {
	p := path.Join(MustMkdirTemp(), socketName)

	unix := fmt.Sprintf("unix://%s", p)

	return must.Do(url.Parse(unix))
}
