# Interfaces

In programming, an interface is a language construct that defines a contract or a set of methods that a class or a type must implement. It specifies the behavior that the implementing class or type should adhere to, without specifying the internal details of how the behavior is implemented.

> An interface defines a set of method signatures without specifying the implementation. It serves as a blueprint for classes or types that implement it.
>
> *Method signatures* : An interface specifies the names, input parameters, output types (if any), and possibly error handling of the methods that implementing types should have.

## Interfaces in go
---
In Go, interfaces play a crucial role in achieving polymorphism and enabling code reusability. They define a set of method signatures that a type must implement to satisfy the interface. 

An interface in Go is defined using the interface keyword followed by a set of method signatures enclosed in curly braces. 

```
    type MyInterface interface {
	    Method1()
	    Method2()
    }
```

In Go, there is no explicit declaration of implementing an interface. A type automatically satisfies an interface if it implements all the methods defined by the interface. This implicit implementation fosters simplicity and flexibility.

**Interface composition** : Go supports interface composition, where an interface can be composed of multiple smaller i*Empty interface*nterfaces. This allows for building larger interfaces by combining smaller reusable interfaces.

 ### Empty Interfaces
 ---
 In Go, the empty interface interface{} does not specify any methods. It represents a type that can hold values of any type. It is commonly used when you want to work with values of different types without knowing their specific types in advance. 

 That means that if we write a function that takes an interface{} value as a parameter, we can supply that function with any value.

 For example, the fmt.Println() function accepts arguments of type interface{} to print values of any type.

 So any function that has it's parameter type `interface{}`, when the function is called with it's argument,Go runtime will perform type conversion (if necessary) and convert the argument value to `interface{}`.

 An interface value is constructed of two words of data; one word is used to point to a method table for the value’s underlying type, and the other word is used to point to the actual data being held by that value



 **References**
   
   1) [Bit of deep dive on interfaces](https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)
   2) [Go Interfaces research paper](https://research.swtch.com/interfaces)



