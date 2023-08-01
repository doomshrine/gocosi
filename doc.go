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

// Package gocosi is a Container Object Storage Interface (COSI) library that simplifies process of writing your own Object Storage Plugin (OSP).
//
// ## Quick Start
//
// The following example illustrates using `gocosi` Object Storage Plugin bootstrapper to create a new COSI Object Storage Plugin from scratch:
//
// ```bash
// mkdir -p cosi-osp && cd cosi-osp
//
//	go run \
//	  github.com/doomshrine/gocosi/cmd/generate@main \
//	  example.com/your/cosi-osp
//
// ```
//
// You will obtain the following file structure in the `cosi-osp` folder:
//
// ```
// cosi-osp
// ├── go.mod
// ├── go.sum
// ├── main.go
// └── servers
//
//	├── identity
//	│   └── identity.go
//	└── provisioner
//	    └── provisioner.go
//
// 4 directories, 5 files
// ```
package gocosi
