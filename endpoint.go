// Copyright © 2023 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gocosi

import (
	"context"
	"errors"
	"io"
	"net"
	"net/url"
	"os"
	"os/user"
	"regexp"
	"strconv"
	"sync"

	"github.com/doomshrine/must"
)

// Schemes supported by COSI.
const (
	SchemeUNIX = "unix"
	SchemeTCP  = "tcp"
)

var digit = regexp.MustCompile(`^\d+$`)

// Endpoint represents COSI Endpoint.
type Endpoint struct {
	user        *user.User
	group       *user.Group
	permissions os.FileMode

	once sync.Once

	address       *url.URL
	listener      net.Listener
	listenerError error
}

// Interface guards.
var _ io.Closer = (*Endpoint)(nil)

// Listener will return listener (and error) after first configuring it.
// Listener is configured only once, on the first call, and the error is
// captured. Every subsequent call will return the same values.
func (e *Endpoint) Listener(ctx context.Context) (net.Listener, error) {
	e.once.Do(func() {
		e.parseEnv() // we parse env here, so user does not need to remember about it.

		listenConfig := net.ListenConfig{}

		if e.address == nil {
			e.listenerError = errors.New("address is empty")
			return
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

// parseEnv is used to call all functions that are picking speicific environment variables.
// The configuration, if valid, will always replace default value.
// Noone of the values parsed here are required to set, so if any of them is invalid,
// the error is logged, if the logger was previously set.
func (e *Endpoint) parseEnv() {
	e.getCOSIEndpoint()
	e.getCOSIEndpointPermissions()
	e.getCOSIEndpointUser()
	e.getCOSIEndpointGroup()
}

// Close ensures that the UNIX socket is removed after the application is closed.
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

func (e *Endpoint) getCOSIEndpoint() {
	env, ok := os.LookupEnv(EnvCOSIEndpoint)
	if !ok || env == "" {
		return
	}

	u, err := url.Parse(env)
	if err != nil {
		log.Error(err, "failed to parse COSI Endpoint from environment")
		return
	}

	log.Info("replacing default", EnvCOSIEndpoint, env)

	e.address = u
}

func (e *Endpoint) getCOSIEndpointUser() {
	env, ok := os.LookupEnv(EnvCOSIEndpointUser)
	if !ok || env == "" {
		return
	}

	var (
		newUser *user.User
		err     error
	)

	if digit.MatchString(env) {
		newUser, err = user.LookupId(env)
	} else {
		newUser, err = user.Lookup(env)
	}

	if err != nil {
		log.Error(err, "failed to parse COSI Endpoint Permissions from environment")
		return
	}

	log.Info("replacing default", EnvCOSIEndpointUser, env)

	e.user = newUser
}

func (e *Endpoint) getCOSIEndpointGroup() {
	env, ok := os.LookupEnv(EnvCOSIEndpointGroup)
	if !ok || env == "" {
		return
	}

	var (
		newGroup *user.Group
		err      error
	)

	if digit.MatchString(env) {
		newGroup, err = user.LookupGroupId(env)
	} else {
		newGroup, err = user.LookupGroup(env)
	}

	if err != nil {
		log.Error(err, "failed to parse COSI Endpoint Permissions from environment")
		return
	}

	log.Info("replacing default", EnvCOSIEndpointGroup, env)

	e.group = newGroup
}

func (e *Endpoint) getCOSIEndpointPermissions() {
	env, ok := os.LookupEnv(EnvCOSIEndpointPerms)
	if !ok || env == "" {
		return
	}

	perms, err := strconv.ParseInt(env, 0, 0)
	if err != nil {
		log.Error(err, "failed to parse COSI Endpoint Permissions from environment")
		return
	}

	log.Info("replacing default", EnvCOSIEndpointPerms, env)

	e.permissions = os.FileMode(perms)
}
