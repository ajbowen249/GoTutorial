![gopher](http://synflood.at/tmp/golang-slides/images/gopher_icon_128.png)
# Go: Is it cool? _Why_ is it cool? Let's find out!
[Go](https://golang.org/) is a modern systems programming language developed primarily by Google with a focus on server development. Its design goals are focused around shedding the weight of old popular systems languages, throwing out things like extesive type hierarchy. Go does have a runtime environment that manages garbage collection, but it is small and compiled into the executable, with no virtual machine or framework to install.

## Categoris to cover (if this is still here, I have failed.)
- Threading
- documentation & examples
- testing
- OOP (how it's different, pointers, anonymous/literal)
- functions DONE
- typecasting DONE!
- for (range)
- short statements
- switch
- defer
- slices/arrays DONE!
- maps DONE!
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
In the first block, we declare a variable of type `interface{}`. We'll talk more about what `interface{}` means later, but for now, just think of it like an "anything goes" type, like `object` in C#. We then assign it the value of 12, and then assign `hiddenInt.(int)` to `exposedInt`. Go does not have true typecasting. What we're doing here is *asserting* the type of `hiddenInt`. If we were have to given `hiddenInt` the value 12.0, the program would crash since that would make `hiddenInt` a `float64`. Later on down, we assign 15.0 to `floatyInt`, and then assign the result of `int(floatyInt)` to `truncatedInt`. Again, this is not casting, but *converting* the `float64` to an `int` with a function. Ass some more code:
```go
    var array [2]int
    array[0] = 12
    array[1] = 32

    fmt.Printf("0: %v, 1: %v\n", array[0], array[1])
```

Here we've made an *array*. *Arrays* in Go look like most other languages. They are always a fixed size, and are always declared with their size in brackets before their type. The more common kind of collections is called a *slice*:
```go
    slice := make([]int, 2)
    slice[0] = 14
    slice[1] = 56

    slice = append(slice, 23)

    fmt.Printf("0: %v, 1: %v, 2: %v\n", slice[0], slice[1], slice[2])
```

*Slice*s look more or less like how you would expect an *array* to look. Behind the scenes, they are basically references to existing arrays that the runtime keeps around, and if you try to append a value beyond their allocated capacity, they will be moved. You can get the length of a slice or array by passing it to the `len()` function, and the capacity of the allocated array for a slice with the `cap()` function. Usage of arrays and slices is a deep subject beyond the scope of this tutorial, though more information is available [here](https://blog.golang.org/go-slices-usage-and-internals).

Go has support for another common data type, the `map`:
```go
    vals := make(map[string]int)
    vals["abc"] = 12
    vals["def"] = 13

    fmt.Printf("abc: %v def: %v\n", vals["abc"], vals["def"])
```

They work more or less as you would expect from other languages. You can also declare literals for *slice*s and `map`s:
```go
    sliceLit := []string{"here", "are", "some", "values"}

    mapLit := map[float64]bool{
        1.54: true,
        2.63: true,
        1.0:  true,
    }

    fmt.Printf("%v\n", sliceLit)
    fmt.Printf("%v\n", mapLit)
```

Add this code:
```go
    value, exists := vals["nonExistantvalue"]
    if exists {
        fmt.Println(value)
    } else {
        fmt.Println("value not found")
    }
```

There are a few new things here. First, note the `if`...`else`. It works like most other languages, and Go does not make you put conditions in parentheses. The bigger deal is on the first line, with the double assignment. Functions in Go can return multiple values. In this case, using the two-value indexing on a `map` gives you the value from that index, as well as a `bool` for whether the value was found. Add this function to your file:
```go
func max2(val1 int, val2 int) (int, bool) {
    if val1 == val2 {
        return val1, true
    } else if val1 > val2 {
        return val1, false
    } else {
        return val2, false
    }
}
```

Erase everything in your existing `main()` function if you want, and call `max2` like this:
```go
    bigger, equal := max2(13, 13)
    if equal {
        fmt.Println("values are equal!")
    } else {
        fmt.Printf("%v is bigger.\n", bigger)
    }
```

`max2` returns the greater of the two values given to it, and a boolean flagging if they were equal. Function declaration syntax is basically `func name(arguments) returnType`. If you have multiple returns, you need the types in parentheses as we have done with `max2`. There are other [advanced aspects](https://tour.golang.org/basics/7) to function declaration. If you want to use `max2` but do not care about the equality flag, try this:
```go
    bigger, _ := max2(13, 13)
    fmt.Printf("%v is bigger.\n", bigger)
```

Underscores in Go act as placeholders for symbol names that syntactically must exist, but are not being used. Go is very picky about unused bindings, and will not compile with them (the same also goes for unused imports).

By now, I bet you're thinking, "cool, no semicolons". You're in for some mild disappointment. Go, like many languages, organizes blocks of code into statements. Go is very picky about its styling, and, for example, will not compile with clamshell-style bracing, only K&R. Technically, having a semicolon everywhere you would expect them is valid Go syntax, but since Go is picky about its styling, it can figure out where statements begin and end without the usual verbosity.