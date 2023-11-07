# go-common

Golang library for common projects on DeBeAndo.

## Usage

To point to the local version of a dependency in Go rather than the one over the web, use the replace keyword.

And now when you compile this module (go install), it will use your local code rather than the other dependency.

```bash
go mod edit -replace github.com/debeando/go-common=$HOME/go/src/github.com/debeando/go-common
```

Revert replacement:

```bash
go mod edit -dropreplace github.com/debeando/go-common
go get -u
```
