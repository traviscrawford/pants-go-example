# pants-go-example
Example Go repo using the pants build system

Most of the code in this repo was copied from
https://github.com/pantsbuild/pants/tree/master/contrib/go/examples to
experiment with different repo layouts, tool integration (e.g.: GoLand),
and other stuff Go users find interesting.

## GoLand

Launching GoLand with the following command appears to work correctly for symbol detection, navigation, etc.
because this command sets `GOPATH` and `GOROOT` per the workspace assembled by Pants.

```
./pants go-env src/go/:: -- goland
```

In the above command, the `::` is a recursive glob for all targets under that directory. All Go targets will be
assembled, along with their 3rdparty dependencies, into a workspace that appears to work with GoLand.

Running main functions worked, as did setting break points and debugging, however, the Pants build cache is
not visible to GoLand, so it does that work again.
