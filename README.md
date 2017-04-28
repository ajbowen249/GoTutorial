![gopher](http://synflood.at/tmp/golang-slides/images/gopher_icon_128.png)
# Go: Is it cool? _Why_ is it cool? Let's find out!
[Go](https://golang.org/) is a modern systems programming language developed primarily by Google with a focus on server development. Its design goals are focused around shedding the weight of old popular systems languages, throwing out things like extesive type hierarchy. Go does have a runtime environment that manages garbage collection, but it is small and compiled into the executable, with no virtual machine or framework to install.

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
In the first block, we declare a variable of type `interface{}`. We'll talk more about what `interface{}` means later, but for now, just think of it like an "anything goes" type, like `object` in C#. We then assign it the value of 12, and then assign `hiddenInt.(int)` to `exposedInt`. Go does not have true typecasting. What we're doing here is *asserting* the type of `hiddenInt`. If we were have to given `hiddenInt` the value 12.0, the program would crash since that would make `hiddenInt` a `float64`. Later on down, we assign 15.0 to `floatyInt`, and then assign the result of `int(floatyInt)` to `truncatedInt`. Again, this is not casting, but *converting* the `float64` to an `int` with a function. Add some more code:
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

Underscores in Go act as placeholders for symbol names that syntactically must exist, but are not being used. Go is very picky about unused bindings, and will not compile with them (the same goes for unused imports).

By now, I bet you're thinking, "cool, no semicolons." You're in for some mild disappointment. Go, like many languages, organizes blocks of code into statements. Go is very picky about its styling, and, for example, will not compile with Allman-style bracing, only K&R. Technically, having a semicolon everywhere you would expect them is valid Go syntax, but since Go is picky about its styling, it can figure out where statements begin and end without the usual verbosity. One place you will still use semicolons (at least, idiomatically) is in *short statements*. To understand them, we'll start with something familiar:
```go
   for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
```
This is pretty self-explanatory, but notice the two semicolons. Like many languages, the `for` loop in Go takes an initialization statement, a conditional statement, and an iterative statement. Since these are all on the same line, the Go compiler still needs the semicolon to delimit them. Why am I explaining this so pedantically? There is an idiomatic pattern in Go that uses *short statement*s in combination with multi-value returns. Here's another way we can call our `max2` function:
```go
    if bigger, equal := max2(13, 13); equal {
        fmt.Println("values are equal!")
    } else {
        fmt.Printf("%v is bigger.\n", bigger)
    }
```

With this, we get the result of `max2` and provide a branching condition on the same line. More importantly, `bigger` and `equal` now only have scope within this chain of `if`...`else` blocks. This is often used to help clean up areas where multiple functions that return a result and an error are used in a row.

Going back to our `for` loop, `for` is actually Go's only looping construct. This will loop forever:
```go
    for {
        fmt.Println("Hellooooo")
    }
```
This is Go's equivalent of a basic while loop:
```go
    i := 0
    for i < 10 {
        i++
    }
```
There are `break` and `continue` keywords that act like those in other languages. Go has its own special kind of automatic `for`:
```go
    vals := []int{1, 5, 2, 9}

    for index, val := range vals {
        fmt.Printf("%v: %v\n", index, val)
    }
```
This will print out the index and value of each item in the iterable. If you don't want the index or don't want the value, replace its name with an underscore. Go also has `switch` statements:
```go
    val := 15
    switch val {
    case 12:
        fmt.Println("It was 12")
    case 15:
        fmt.Println("It was 15")
    default:
        fmt.Println("It was something")
    }
```
Notice that there is no `break` statement used here. Go flips the logic of the traditional `switch`-`case` structure by automatically breaking at the end of a `case` block, and requiring the `fallthrough` keyword to explicitly fall through to the next case:
```go
    val := 3
    switch val {
    case 6:
        fmt.Println("It was bigger than 5")
        fallthrough
    case 5:
        fmt.Println("It was bigger than 4")
        fallthrough
    case 4:
        fmt.Println("It was bigger than 3")
        fallthrough
    case 3:
        fmt.Println("It was bigger than 2")
        fallthrough
    case 2:
        fmt.Println("It was bigger than 1")
        fallthrough
    default:
        fmt.Println("It was bigger than 0")
    }
```
You can `switch` without a variable with a set of boolean expression `case`s as an alternative to a big `if`...`else` but with the ability to `fallthrough`:
```go
    val := 3
    switch {
    case val == 1:
        fmt.Println("val is 1")
    case val > 2:
        fmt.Println("val is more than 2")
        fallthrough
    case val > 1:
        fmt.Println("val is more than 1")
        fallthrough
    case val > 0:
        fmt.Println("val is more than 0")
    case val == 9:
        fmt.Println("val is 9")
    default:
        fmt.Println("val is something")
    }
```
Note that `case`s in this construct are evaluated in order, so if I were to have a `case val == 3:` at the end, it would not be hit, but it would if it was before the `case val > 2:`.

Functions in Go are first-class. Here's a whole new main file:
```go
package main

import "fmt"

func main() {
    doCompare(12, 1, greater, printIfTrue)
}

func greater(x int, y int) bool {
    return x > y
}

func printIfTrue(value bool) {
    if value {
        fmt.Println("true")
    } else {
        fmt.Println("not true")
    }
}

func doCompare(x int, y int, compare func(int, int) bool, act func(bool)) {
    act(compare(x, y))
}

```
Function type declarations are just like function signature lines, but without names. You can also assign functions to variables:
```go
    printIfFalse := func(value bool) {
        if !value {
            fmt.Println("false")
        } else {
            fmt.Println("not false")
        }
    }

    doCompare(12, 1, greater, printIfFalse)
```
Which means you can use functions anonymously:
```go
    doCompare(12, 12, func(x int, y int) bool {
        return x == y
    }, func(equal bool) {
        if equal {
            fmt.Println("equal")
        } else {
            fmt.Println("not equal")
        }
    })
```
Functions as declared in the previous two examples can capture variables in a closure like other languages. Go has lightweight Object-Oriented Programming concepts:
```go
package main

import "fmt"

func main() {
    point1 := point{0, 0}
    fmt.Println(point1.ToString())

    point1.TranslateX(5)
    fmt.Println(point1.ToString())

    point1.TranslateY(-3)
    fmt.Println(point1.ToString())
}

type point struct {
    x int
    y int
}

func (p point) ToString() string {
    return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func (p point) TranslateX(delta int) {
    p.x += delta
}

func (p point) TranslateY(delta int) {
    p.y += delta
}
```
Struct methods are implemented by adding an argument taking an instance of the particular Struct before the function name. You can implement these functions anywhere you like, so you can even extend `struct`s from the standard or third-party libraries. When we instantiated our `point`, we supplied values in the order that they are declared in the `struct` itself. We can also name them explicitly:
```go
    point1 := point{x: 2, y: 5}
```
We can also omit values and get the type's defaults:
```go
    point1 := point{y: 6}
    point2 := point{}
```
Here's where things start to get weird: Go's does not have keywords like "public" and "private". Go has different concepts of *exported* and *unexported*. Within a `package`, Any function, `struct` method, or `struct` field whose name starts with a capitol letter is considered *exported*, and is visible outside of the package. Otherwise, they are *unexported*, and are only usable from within a package. This can be useful when you write unit tests, as `struct` fields can be examined from a test and behavior can be white-box-tested if you're into that sort of thing. This has two main drawbacks to watch out for:
- Design patterns like having accessor methods for property fields are still encouraged, but difficult to enforce from within a package.
- Forgetting to name things against this convention can lead to headaches when you need to refactor a name just to use it in another package.

Go does not have polymorphism, but it does have `interface`s:
```go
type stringable interface {
    ToString() string
}

func Print(val stringable) {
    fmt.Println(val.ToString())
}
```
We can now call `Print` on an instance of our `point`, without having to add anything to the declaration of `point`. If a type doesn't have a function it needs to satisfy at type assertion, the compiler will complain. Since the system is so lightweight, `interface`s can be declared anonymously:
```go
func Print(val interface {
    ToString() string
}) {
    fmt.Println(val.ToString())
}
```
It should now be clear what we meant earlier by `interface{}`. What we meant was, "this is a variable that implements this anonymous interface," and the `interface` we declared there was empty, so anything will satisfy it.

Go has a fun and unique threading model. Let's start with something familiar (note that this is NOT idiomatic Go code):
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    mutex := sync.Mutex{}
    i := 0

    lock := func(action func()) {
        mutex.Lock()
        action()
        mutex.Unlock()
    }

    done := func() bool {
        var done bool
        lock(func() { done = i >= 10 })
        return done
    }

    go func() {
        for !done() {
            lock(func() { fmt.Println(i) })
        }
    }()

    go func() {
        for !done() {
            lock(func() { i++ })
        }
    }()

    for !done() {
        time.Sleep(10 * time.Millisecond)
    }

    fmt.Println("Done!")
}
```
The first thing you'll notice is the `go` keyword. Threads in Go are called *goroutines*. Unlike many languages, You don't get handles to threads to join on. To run a new thread, all you need to do is `go someFunction()`. *Goroutine*s always operate on a function call, which is why we need the extra `()` at the end of the function literals in the example above. When we examine what this program does, we see that we have two threads operating on the same `i` variable, so we mutex around accessing it as you would expect, though I did make a special `lock` function for convenience. When we run this, it will print each value of `i` a non-deterministic number of times from 0 to 9. Let's re-write this using a new mechanism:
```go
package main

import "fmt"

func main() {
    value := make(chan int)
    go produce(10, value)
    consume(9, value)
}

func produce(num int, value chan int) {
    for i := 0; i < num; i++ {
        value <- i
    }
}

func consume(target int, value chan int) {
    for {
        i := <-value
        fmt.Println(i)
        if i == target {
            break
        }
    }
}

```
There are a few improvements here. First, running the program will only print once per new value. More importantly, we're now passing data through a special construct called a `chan`nel. `chan`nels can be thought of as a sort of "pipe" between threads that allow for transport of specific data. If you look at our `produce` function, you can see that we send `i` through the `value` chan using the `<-` operator. In `consume`, we receive a value from the `value` chan with the same operator. Think of `<-` as an arrow that points where data is going. By default, sending data will block until a receiver has read the previous value, and receiving will block until a value is present.

If we had declared our `chan`nel with `value := make(chan int, 5)`, we would have a *buffered `chan`nel*. In this case, `produce` will not block writing so long as the `chan`nel has less than 5 values in its buffer. Reading still works the same way as an unbuffered `chan`nel, but with the chance that a value will be immediately available in some iterations.

There's an even simpler way to write this program:
```go
package main

import "fmt"

func main() {
    value := make(chan int)
    go produce(10, value)
    consume(value)
}

func produce(num int, value chan int) {
    for i := 0; i < num; i++ {
        value <- i
    }

    close(value)
}

func consume(value chan int) {
    for i := range value {
        fmt.Println(i)
    }
}
```
The big differences here are the `close` call in `produce`, and the simpler loop in `consume`. With this mechanism, a receiver of data from a `chan`nel can use the `range` keyword to loop over it like an iterable, and the loop will exit when the `chan`nel has been `close`d. Similar to the `map`, you can also assign the result of a `chan`nel read to two variables, like `value, open := <- channel`, and get a flag signaling that the `chan`nel is still open. There are some other [advanced topics](https://tour.golang.org/concurrency/5) with `chan`nels to discover.

There's another way we could have `close`d our `chan`nel:
```go
func produce(num int, value chan int) {
    defer close(value)

    for i := 0; i < num; i++ {
        value <- i
    }
}
```

The `defer` keyword stacks up a function call to execute after the function has returned, and will still exit even if the function `panic`s. It's often used for resource cleanup. If you have multiple `defer`s in a function, their operations will stack, not queue.

# Go In Production
I've mentioned `package`s a few times now. `package`s are Go's concept of libraries. The current state of Go's dependency management is somewhat infantile, buf if you see a package you like online, you can `go get` it from the command line. You can even get this tutorial as a package. How about you do that, and meet me back on the other side:
```
go get github.com/ajbowen249/GoTutorial
```
When that completes, you should have the code for this repository in `$GOPATH/src/github.com/ajbowen249/GoTutorial`. If you `go install`, you'll get an executable at `$GOPATH/bin/GoTutorial` that does some basic benchmarking of two sorting algorithms and prints out the data. The app is comprised of code lazily copied in and slightly tweaked from my [GoSandbox](https://github.com/ajbowen249/GoSandbox) repo.

Examining this project, we have this structure:
```
+GoTutorial
├—+algorithms
│ ├—algorithms.go
│ └—algorithms_test.go
│
├—+datastructures
│ ├—binaryTree.go
│ ├—linkedList.go
│ └—datastructures_test.go
│
├—+table
│ ├—table.go
│ └—table_test.go
│
└—main.go
```
The primary `package` here is `GoTutorial`, and the `algorithms`, `datastructures`, and `table` packages are all relative to it. Each package has a complimentary "_test.go" file. Go has a built-in system for running unit tests. If you wish to run the tests in this package, run `go test ./...` from the command line (note that there are three dots after the slash). The basics of Go's testing library are fairly self-explanatory if you examine the test files, and more information is available [here](https://www.golang-book.com/books/intro/12).

If you take a look at main.go, you'll see this near the top:
```go
import (
    "fmt"
    al "github.com/ajbowen249/GoTutorial/algorithms"
    ds "github.com/ajbowen249/GoTutorial/datastructures"
    "github.com/ajbowen249/GoTutorial/table"
    "math/rand"
    "time"
)
```
There are two new things here. First, note that the fully-qualified names of the packages also include "github.com" and my handle. If you started from scratch, you could find these packages in your `$GOPATH/src` with their fully-qualified name as the directory. Second, note that I give the aliased names `al` and `ds` to `algorithms` and `datastructures`, respectively. This is mainly for convenience later on.

If you want to see documentation on this package, use `go doc` in your command line. You won't see much just running that from the GoTutorial root, so try running `go doc algorithms`. This will print out a high-level overview of the `algorithms` package. `go doc algorithms.SortInt` will expand on the documentation of the `SortInt` function. The high-level overview will list all of the *exported* functions in the package, so anything that doesn't start with a capitol letter will not be seen here. Take a look at the algorithms/algorithms.go file. You'll see special comments above each function. Your Go linter will be very picky about how these are worded, and they are what shows up in the result of the `go doc` tool.

There is also a far prettier way to see this documentation. try running `godoc -http=:8080`. Note that the command here is `godoc`, not `go doc` (confused? [you're not alone](http://whipperstacker.com/2015/09/30/go-documentation-godoc-godoc-godoc-org-and-go-doc/)). Now open up [http://localhost:8080/pkg/](http://localhost:8080/pkg/). `godoc` is running a complete server hosting an offline version of the Go documentation, and if you look through the packages section, you'll see the nicely-formatted documentation of the standard library alongside every package you have in your `$GOPATH`, even your own code!

## Resources
That's it! Here are a few links for some more information, should you wish to look:
- [The Go homepage](https://golang.org/)
- [A Tour of Go](https://tour.golang.org/welcome/1)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [More Go Styling and Convention](https://github.com/golang/go/wiki/CodeReviewComments)
