# go-common

Golang library for common projects on DeBeAndo.

## Who work it?

If you want to say, point to the local version of a dependency in Go rather than the one over the web, use the replace keyword.

And now when you compile this module (go install), it will use your local code rather than the other dependency.

```bash
$ go mod edit -replace github.com/debeando/go-common=/Users/nsc/go/src/github.com/debeando/demo
```

Following the -replace is first what you want to replace, then an equals sign, then what you’re replacing it with.
