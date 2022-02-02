# HTTP Server

[![Build Test](https://github.com/matteeyao/echo-server/actions/workflows/build-test.yml/badge.svg)](https://github.com/matteeyao/echo-server/actions/workflows/build-test.yml)

From a command prompt, run the `main.go` file using the command:

```zsh
$ go run .
```

or

```zsh
$ go run .
```

To determine whether the unit tests are passing or failing run:

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

## HTTP Server Specs

The specifications of this HTTP server are covered in the integration tests within the `http_server_spec` submodule

1. In the root directory run:

```zsh
go run .
```

This will run the server on port 5000

2. Next, `cd` into the `http_server_spec` directory:

```zsh
cd http_server_spec
```

3. Once the server is running, you can run the acceptance test suite with:

```zsh
rake test
```

You can also run the tests from a specific section of the features:

```zsh
rake test:f1 # Run all of the tests in 01_getting_started
rake test:f2 # Run all of the tests in 02_structured_data
rake test:f3 # Run all of the tests in 03_file_server
rake test:f4 # Run all of the tests in 04_todo_list
```

## Dependencies

To complete this tutorial, you'll need the following:

* [Go Version 1.17](https://golang.org/dl/) or higher installed on your local machine. You can follow these instructions to install Go on [Linux](https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-ubuntu-18-04), [macOS](https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-macos) and [Windows](https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-windows-10). On macOS, you can also install Go using the [Homebrew package manager](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-homebrew-on-macos).

## Resources

* [Build Web Application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/)

* [The Go Programming Language](https://learning.oreilly.com/library/view/the-go-programming/9780134190570/) By Alan A. A. Donovan, Brian W. Kernighan

* [Setting up Github actions](https://medium.com/swlh/setting-up-github-actions-for-go-project-ea84f4ed3a40)
