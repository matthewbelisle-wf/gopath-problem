gopath-problem
==============

This repo explains a problem related to goapp.  The problem is that you cannot have multiple
packages that are testable and depend on each other in the same github repo.  The reason is because
`goapp test` sets `$GOPATH` to the directory of the `*_test.go` as it runs the test, which means
that importing local package that were originally in `$GOPATH` during `goapp serve` are no longer
in the `$GOPATH` and you get an import error.

```
# serve runs fine
$ goapp serve .
...
# test breaks
$ goapp test .
main.go:5:2: cannot find package "bar" in any of:
        /Users/mgbelisle/gae/go_appengine/goroot/src/pkg/bar (from $GOROOT)
        /Users/mgbelisle/go/src/bar (from $GOPATH)
$ goapp test ./bar/
# test breaks on the individual packages too
bar/bar.go:3:8: cannot find package "foo" in any of:
        /Users/mgbelisle/gae/go_appengine/goroot/src/pkg/foo (from $GOROOT)
        /Users/mgbelisle/go/src/foo (from $GOPATH)
```

The same problem is described
[here](https://groups.google.com/forum/#!topic/google-appengine-go/C_i5kQEi7-A).  The solution is
to use one repo per package.