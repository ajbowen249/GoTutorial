![gopher](http://synflood.at/tmp/golang-slides/images/gopher_icon_128.png)
# Go: Is it cool? _Why_ is it cool? Let's find out!
[Go](https://golang.org/) is a modern systems programming language developed primarily by Google with a focus on server development. Its design goals are focused around shedding the weight of old popular systems languages, throwing out things like extesive type hierarchy. Go does have a runtime environment that manages garbage collection, but it is small and compiled into the executable, with no virtual machine or framework to install.

## Categoris to cover (if this is still here, I have failed.)
- Threading
- documentation & examples
- testing
- OOP (how it's different, pointers, anonymous/literal)
- functions
- typecasting
- for (range)
- short statements
- switch
- defer
- slices/arrays
- maps
- closures

## Let's Get Going
If you haven't already, run through the [setup and install](https://golang.org/doc/install). By this point, you should have something like:
```go
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```

When you did a `go install` or a `go build` for this file, it produced a native executable with the name of the containing folder of the file. To make a package into an executable in Go, you need a file in your package that exists in the `main` package, and has a `main()` function. We'll talk more about packages later.

For now, note the `import "fmt"` and `fmt.Printf("hello, world\n")` lines.