![gopher](http://synflood.at/tmp/golang-slides/images/gopher_icon_128.png)
# Go: Is it cool? _Why_ is it cool? Let's find out!
[Go](https://golang.org/) is a modern systems programming language developed primarily by Google with a focus on server development. Its design goals are focused around shedding the weight of old popular systems languages, throwing out things like extesive type hierarchy. Go does have a runtime environment that manages garbage collection, but it is small and compiled into the executable, with no virtual machine or framework to install.

## Categoris to cover (if this is still here, I have failed.)
- Threading
- documentation & examples
- testing
- OOP (how it's different, pointers, anonymous/literal)
- functions
- typecasting DONE!
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
    fmt.Println("hello, world")
}
```

When you did a `go install` or a `go build` for this file, it produced a native executable with the name of the containing folder of the file. To make a package into an executable in Go, you need a file in your package that exists in the `main` package, and has a `main()` function. We'll talk more about packages later.

For now, note the `import "fmt"` and `fmt.Println("hello, world\n")` lines. `fmt` is a package from the Go standard library focused on formatting and writing strings. `Println` prints a string to standard output with a newline. Nothing shocking here.

Let's go over basic variables. Add the following to your `main()` function:
```go
    var x int
    x = 3

    y := 1.23

    z := false

    fmt.Printf("x: %v y: %v z: %v\n", x, y, z)
```
On the first line, we declare a `var`iable `x` of type `int`. At this point, it will be initialized with the default `int` value, which is 0. We then assign the value 3 to `x`. Later on, we create a float variable, `y`, in one line. When you want to create and initialize a variable all in one line, use the `:=` operator. You don't need to declare a type here, because the compiler can infer it based on the right-hand side. When using literal numeric values, the compiler will assume they are `int` or `float64`. Go's primitive types are `bool`, `string`, `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`, `byte`, `rune`, `float32`, `float64`, `complex64`, and `complex128`. Special notes:
- `int`, `uint`, and `uintptr` all alias to their 32 or 64-bit counterparts based on the architecture of the system.
- `rune` is a unicode char point in a string.
- `complex64` and `complex128` represent complex numbers.
On the last line we added, we use a format string. The `fmt.Printf` function takes a string and a variadic set of arguments to insert into that string in place of the format fields. `%v` in this case represents the "value" format type, and is probably the most common one you'll use. You can find more detail on other format types [here](https://gobyexample.com/string-formatting). Add some more code:
```go
    var hiddenInt interface{}
    hiddenInt = 12
    exposedInt := hiddenInt.(int)

    floatyInt := 15.0
    truncatedInt := int(floatyInt)

    fmt.Printf("exposedInt: %v truncatedInt: %v\n", exposedInt, truncatedInt)
```
In the first block, we declare a variable of type `interface{}`. We'll talk more about what `interface{}` means later, but for now, just think of it like an "anything goes" type, like `object` in C#. We then assign it the value of 12, and then assign `hiddenInt.(int)` to `exposedInt`. Go does not have true typecasting. What we're doing here is *asserting* the type of `hiddenInt`. If we were have to given `hiddenInt` the value 12.0, the program would crash since that would make `hiddenInt` a `float64`. Later on down, we assign 15.0 to `floatyInt`, and then assign the result of `int(floatyInt)` to `truncatedInt`. Again, this is not casting, but *converting* the `float64` to an `int` with a function.