
# Functions

Go functions requires explicit returns, i.e. it won’t automatically return the value of the last expression.


### Init function
---
In Go, the init function is a special function that can be defined in any package. The init function is executed automatically before the main function is executed. Its purpose is to perform package initialization tasks or setup actions that need to be executed before the program starts.

 The init function has no parameters and no return type. It is defined as func init().

  The init function is automatically executed once when the program starts, before the main function is called.

   The init function is often used for tasks such as initializing variables, setting up database connections, registering drivers, configuring logging, or performing any other package-level setup actions.

### main function
---
Every piece of Go code is written inside a main package and the package that begins a Go program has a function called main().This function is the entry point of Go program execution.

```
func functionName(parameters) returnType {
    // Function body
    // Code to be executed
    // Return statement (optional)
}
```

### Closure function
---
In Go we can define a function within a function. This type of function is called a closure function.

Closure functions are anonymous and have almost the same syntax as normal functions except that they have no name.

In Go, a closure is a function value that references variables from outside its body. The closure function has access to the variables in its surrounding lexical scope, even after the outer function has finished execution. It "closes over" the variables it references, hence the term "closure."

> Closures are useful in scenarios where you want to create functions that encapsulate some state or behavior. They allow you to define functions that have access to variables or data specific to their creation context. This can be particularly useful for creating callback functions, managing stateful operations, or implementing function factories.


