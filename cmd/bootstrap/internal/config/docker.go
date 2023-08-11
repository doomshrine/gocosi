// Copyright © 2023 gocosi authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

const (
	DefaultRootlessImage = "docker.io/rockylinux/rockylinux:9-ubi"
	DefaultImage         = "gcr.io/distroless/static:latest"
)

type Docker struct {
	RootlessImage string
	Image         string
}

func newDocker() *Docker {
	return &Docker{
		RootlessImage: DefaultRootlessImage,
		Image:         DefaultImage,
	}
}
