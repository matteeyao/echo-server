# Echo Server

[![Build Test](https://github.com/matteeyao/echo-server/actions/workflows/build-test.yml/badge.svg)](https://github.com/matteeyao/echo-server/actions/workflows/build-test.yml)

From a command prompt, run the `main.go` file

```zsh
$ go run main.go
```

To determine whether the tests are passing or failing run:

```zsh
$ go test
```

For a more detailed test output:

```zsh
$ go test -v
```

Run the following command to calculate the coverage for your current unit tests:

```zsh
$ go test -coverprofile=coverage.out
```

## Dependencies

To complete this tutorial, you'll need the following:

* [Go Version 1.17](https://golang.org/dl/) or higher installed on your local machine. You can follow these instructions to install Go on [Linux](https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-ubuntu-18-04), [macOS](https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-macos) and [Windows](https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-windows-10). On macOS, you can also install Go using the [Homebrew package manager](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-homebrew-on-macos).

## Resources

* [The Go Programming Language](https://learning.oreilly.com/library/view/the-go-programming/9780134190570/) By Alan A. A. Donovan, Brian W. Kernighan

* [Setting up Github actions](https://medium.com/swlh/setting-up-github-actions-for-go-project-ea84f4ed3a40)
