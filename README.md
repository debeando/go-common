# go-common

Golang library for common projects on DeBeAndo.

## Who work it?

To point to the local version of a dependency in Go rather than the one over the web, use the replace keyword.

And now when you compile this module (go install), it will use your local code rather than the other dependency.

```bash
go mod edit -replace github.com/debeando/go-common=/Users/nsc/go/src/github.com/debeando/demo
```
