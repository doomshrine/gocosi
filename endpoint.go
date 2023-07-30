package gocosi

import (
	"context"
	"errors"
	"net"
	"net/url"
	"os"
	"os/user"
	"strconv"
	"sync"

	"github.com/doomshrine/must"
)

const (
	SchemeUNIX = "unix"
	SchemeTCP  = "tcp"
)

type Endpoint struct {
	user        *user.User
	group       *user.Group
	permissions os.FileMode

	once sync.Once

	address       *url.URL
	listener      net.Listener
	listenerError error
}

func (e *Endpoint) Listener(ctx context.Context) (net.Listener, error) {
	e.once.Do(func() {
		listenConfig := net.ListenConfig{}

		if e.address == nil {
			e.listenerError = errors.New("address is empty")
		}

		e.listener, e.listenerError = listenConfig.Listen(ctx, e.address.Scheme, e.address.Path)
		if e.listenerError != nil {
			return
		}

		e.listenerError = e.chown()
		if e.listenerError != nil {
			return
		}

		e.listenerError = e.setPermissions()
		if e.listenerError != nil {
			return
		}
	})

	return e.listener, e.listenerError
}

func (e *Endpoint) Close() error {
	if e.address != nil && e.address.Scheme == SchemeUNIX {
		return os.Remove(e.address.Path)
	}

	return nil
}

func (e *Endpoint) chown() error {
	if e.user == nil || e.group == nil || e.address.Scheme != SchemeUNIX {
		return nil
	}

	uid := must.Do(strconv.Atoi(e.user.Uid))
	gid := must.Do(strconv.Atoi(e.group.Gid))

	return os.Chown(e.address.Path, uid, gid)
}

func (e *Endpoint) setPermissions() error {
	if e.address.Scheme != SchemeUNIX {
		return nil
	}

	return os.Chmod(e.address.Path, e.permissions)
}
