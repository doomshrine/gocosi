# Contributing to gocosi

**Table of Contents**

- [Contributing to gocosi](#contributing-to-gocosi)
  - [Issues](#issues)
    - [Reporting an Issue](#reporting-an-issue)
    - [Issue Lifecycle](#issue-lifecycle)
  - [Pull Requests](#pull-requests)
  - [Developing](#developing)
  - [Developing](#developing-1)
    - [Go Environment and Go Modules](#go-environment-and-go-modules)
    - [Code Linting with golangci-lint](#code-linting-with-golangci-lint)
      - [Installing golangci-lint via Homebrew (macOS)](#installing-golangci-lint-via-homebrew-macos)
      - [Installing golangci-lint via `go install`](#installing-golangci-lint-via-go-install)
    - [Testing](#testing)
      - [Writing Tests](#writing-tests)
        - [Unit tests](#unit-tests)
        - [Integration tests](#integration-tests)
      - [Testify Library](#testify-library)
  - [Releasing](#releasing)
    - [Tagging a release](#tagging-a-release)
      - [Prerequisites](#prerequisites)
      - [Tagging the release](#tagging-the-release)
    - [If a release fails](#if-a-release-fails)
      - [Github Releases](#github-releases)
        - [Prerequisites](#prerequisites-1)

**First:** if you're unsure or afraid of _anything_, just ask
or submit the issue or pull request anyways. You won't be yelled at for
giving your best effort. The worst that can happen is that you'll be
politely asked to change something. We appreciate all contributions!

For those folks who want a bit more guidance on the best way to
contribute to the project, read on. Addressing the points below
lets us merge or address your contributions quickly.

## Issues

### Reporting an Issue

* Make sure you test against the latest released version. It is possible
  we already fixed the bug you're experiencing.
* If you experienced a panic, please create a [gist](https://gist.github.com)
  of the *entire* generated crash log for us to look at. Double check
  no sensitive items were in the log.
* Respond as promptly as possible to any questions made by the _gocosi_
  team to your issue. Stale issues will be closed.

### Issue Lifecycle

1. The issue is reported.
2. The issue is verified and categorized by a _gocosi_ collaborator.
   Categorization is done via labels. For example, bugs are marked as "bugs".
3. Unless it is critical, the issue is left for a period of time (sometimes
   many weeks), giving outside contributors a chance to address the issue.
4. The issue is addressed in a pull request. The issue will be
   referenced in commit message(s) so that the code that fixes it is clearly
   linked.
5. The issue is closed. Sometimes, valid issues will be closed to keep
   the issue tracker clean. The issue is still indexed and available for
   future viewers, or can be re-opened if necessary.

## Pull Requests

Pull requests must always be opened from a fork of `gocosi`, even if you have
commit rights to the repository so that all contributors follow the same process.

## Developing

## Developing

### Go Environment and Go Modules

To contribute to gocosi, you need to have Go installed on your system and set up with Go modules. Follow these steps to get started:

1. Install Go:
   - For macOS users, the recommended way is to use Homebrew:
     ```
     $ brew install go
     ```
   - For other platforms or manual installation, you can download and install Go from the [official website](https://golang.org/dl/).

2. Clone the `gocosi` repository to your local machine:
   ```
   $ git clone https://github.com/$YOUR_USERNAME/gocosi.git
   ```

3. Change into the `gocosi` directory:
   ```
   $ cd gocosi
   ```

4. Now you're all set with the Go environment and Go modules!

### Code Linting with golangci-lint

To ensure consistent code quality, we use `golangci-lint` as a single point for code linting. You can install `golangci-lint` via Homebrew (for macOS users) or using the `go install` command (for all platforms).

#### Installing golangci-lint via Homebrew (macOS)

If you're on macOS and using Homebrew, you can install `golangci-lint` with the following command:

```sh
$ brew install golangci-lint
```

Make sure to update `golangci-lint` regularly to get the latest improvements:

```sh
$ brew upgrade golangci-lint
```

#### Installing golangci-lint via `go install`

For other platforms, you can install `golangci-lint` using the `go install` command:

```sh
$ go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

Ensure that your Go binary is in your system's PATH for the `go install` command to work correctly.

With `golangci-lint` installed, you can now run it against the gocosi codebase to check for any linting issues:

```sh
$ golangci-lint run
```

Fix any linting issues reported by `golangci-lint` before submitting your changes.

Remember, we encourage contributions to be well-formatted and follow the project's coding conventions. Happy coding!

### Testing

#### Writing Tests

When adding new features or fixing bugs, it's essential to write tests to ensure the stability and correctness of the code changes. `gocosi` uses both unit tests and integration tests.

##### Unit tests

Unit tests focus on testing individual functions and components in isolation. To write a unit test, create a new file in the `*_test.go` format alongside the code you want to test. Use the Go testing framework along with the testify/assert and testify/require libraries to create test functions that cover different scenarios and edge cases.

Example unit test using testify/assert:

```go
import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    assert.Equal(t, 5, result, "Expected the sum of 2 and 3 to be 5")
}
```

##### Integration tests

Integration tests check the interaction between different parts of the system and may involve external dependencies like databases or APIs. Write integration tests in separate test files with appropriate names to differentiate them from unit tests.

Example integration test using testify/require:

```go
import (
    "testing"
    "github.com/stretchr/testify/require"
)

func TestIntegrationFlow(t *testing.T) {
    // Set up test environment and dependencies
    // Run the integration flow
    // Check the expected output using require/assert functions
}
```

#### Testify Library

The testify library provides a set of utilities for test assertions and mocking in Go. It includes the `assert` package for simple assertions and the `require` package for assertions that stop the test immediately upon failure. Additionally, the `mock` package can be used for creating mock objects during testing.

Make sure to use these libraries effectively to write comprehensive tests and ensure the reliability of the `gocosi` codebase.

## Releasing

### Tagging a release

When it's time to make a new release of `gocosi`, follow these steps to tag the release:

#### Prerequisites

Before tagging a release, make sure:
- All changes for the release are merged into the `main` branch.


#### Tagging the release

1. Create an annotated tag for the release:

```sh
$ git tag -a v1.0.0 -m "Release version 1.0.0"
```

2. Push the tag to the remote repository:

```sh
$ git push origin v1.0.0
```

3. Github Actions will automatically trigger the release workflow, which will create a new release on the repository with the tag information.

### If a release fails

If a release fails for any reason, follow these steps to handle the situation:

#### Github Releases

##### Prerequisites

Before attempting to create a new release, make sure:
- You have the necessary permissions to create a release on the repository.
- Check the release workflow and ensure it's properly configured to handle the release process.

If the release workflow fails:
1. Investigate the cause of the failure by reviewing the logs and error messages.
2. Make necessary fixes to the release workflow or the repository configuration.
3. Retry the release workflow.

Remember, creating a new release is a critical process, so always double-check everything before proceeding.

Remember, this document is a starting point for contributors to understand how to work with gocosi and contribute effectively. It's important to keep it up to date and include any changes in the development and contribution processes over time.

Feel free to extend and modify this document to reflect any new practices or guidelines for contributing to gocosi. Happy contributing!