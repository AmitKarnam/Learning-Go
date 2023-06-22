# Constants

Go supports constants of character, string, boolean, and numeric values.

In Go, once a constant is declared and assigned a value, its value cannot be changed or reassigned. Constants are immutable.

> `const` declares a constant value.

Constant expressions perform arithmetic with arbitrary precision.
A numeric constant has no type until it’s given one, such as by an explicit conversion.

*Go is a statically typed language that does not permit operations that mix numeric types. You can’t add a float64 to an int, or even an int32 to an int.*

---


## Background : C

In the early days of thinking about Go, the makers talked about a number of problems caused by the way C and its descendants that let you mix and match numeric types. Many mysterious bugs, crashes, and portability problems are caused by expressions that combine integers of different sizes and “signedness”.

When designing Go, The makers decided to avoid this minefield by mandating that there is no mixing of numeric types.

> `const hello = "Hello"`

This is an untyped string constant, which is to say it is a constant textual value that does not yet have a fixed type. Yes, it’s a string, but it’s not a Go value of type string. It remains an untyped string constant even when given a name.

`hello ` is also an untyped string constant. An untyped constant is just a value, one not yet given a defined type that would force it to obey the strict rules that prevent combining differently typed values.

So what, then, is a typed string constant? It’s one that’s been given a type, like this :

> `const typedHello string = "Hello, World!!"`

<br/>

Ref :  [Go Constants](https://go.dev/blog/constants)




 