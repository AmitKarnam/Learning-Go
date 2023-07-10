# Methods

We do not have classes in go programming language. We can make illusion of classes by *defining functions on types*. This type is called ***receiver argument***.

Receiver argument is statement between *func* keyword and function name in declaration. Function defined on type is called method.

Each variable of receiver argument type declartion in package can call defined method.

***Receiver type must be in the same package as method.***

```
type Rectangle struct {
    a,b int
}

func (r *Rectangle) Area() int {
    return r.a * r.b
}

func main() {
    r := Rectangle(2,3)
    fmt.Println(r.Area())
}

```

The function *Area()* is a method on *Rectangle* type.

### Receiver Type
---

Receiver type can of 2 major types : 

- Value Receiver
- Pointer Receiver

**Value Receiver**
1. Syntax: `(r ReceiverType)`
2. In a value receiver, the method operates on a copy of the original instance of the type.
3. Any modifications made within the method do not affect the original instance.
4. Value receivers are useful when you want to ensure that the original instance remains unchanged during the method execution. 

```
type Rectangle struct {
    width  float64
    height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func main() {
    rect := Rectangle{width: 5, height: 3}
    area := rect.Area()
    fmt.Println("Area:", area)
}

```

**Pointer Receiver**
1. Syntax: `(r *ReceiverType)`
2. In a pointer receiver, the method operates directly on the original instance of the type by using a pointer to that instance.
3. Modifications made within the method can affect the original instance.
4. Pointer receivers are useful when you want to modify the original instance or avoid copying large instances of the type.

```
type Rectangle struct {
    width  float64
    height float64
}

func (r *Rectangle) Scale(factor float64) {
    r.width *= factor
    r.height *= factor
}

func main() {
    rect := &Rectangle{width: 5, height: 3}
    rect.Scale(2)
    fmt.Println("Width:", rect.width, "Height:", rect.height)
}
```


> Choosing between a value receiver and a pointer receiver depends on our requirements. If we need to modify the instance or if the type is large, a pointer receiver is typically used. On the other hand, if we want to operate on a copy of the instance or if the type is small, a value receiver can be used.